package dfabric

import (
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric/common/tools/cryptogen/ca"
	"github.com/hyperledger/fabric/common/tools/cryptogen/msp"
	"github.com/pkg/errors"
)

func BuildPeerNodeByFiles(output, OrgPath, baseDomain, host, user string) (err error) {
	signca, tlsca, err := LoadCAObjectFromFiles(OrgPath, baseDomain)
	if err != nil {
		return err
	}

	err = BuildPeerNode(output, signca, tlsca, baseDomain, host, user)
	if err != nil {
		return err
	}

	err = copyAdminCert(filepath.Join(OrgPath, "users"), filepath.Join(output, "msp", "admincerts"), "Admin@"+baseDomain)
	if err != nil {
		err = errors.WithMessage(err, "copy user cert error")
		return
	}
	return nil
}

func BuildPeerNode(output string, signca, tlsca *ca.CA, baseDomain, host, user string) (err error) {
	url := host + "." + baseDomain

	err = msp.GenerateLocalMSP(filepath.Join(output, "peer"), url, []string{url}, signca, tlsca, msp.PEER, false)
	if err != nil {
		err = errors.WithMessage(err, "build MSP error")
		return
	}

	err = msp.GenerateLocalMSP(filepath.Join(output, "user"), user+"@"+baseDomain, []string{url}, signca, tlsca, msp.CLIENT, false)
	if err != nil {
		err = errors.WithMessage(err, "build MSP error")
		return
	}
	return nil
}

func BuildNewPeerInMSP(OrgPath, baseDomain, host, user string) (err error) {
	// OrgPath /---
	//           |- peers /- host.baseDomain: should not be exist
	//           |- users /- user@baseDomain: should not be exist
	//           |- ca : should be exist
	//           |- tlsca : should be exist

	// 1. check ca files
	signCA, tlsCA, err := LoadCAObjectFromFiles(OrgPath, baseDomain)
	if err != nil {
		return err
	}

	// 2. check peer files
	// 3. make peer files
	if len(host) > 0 {
		err = subBuildNewPeerInMSP(signCA, tlsCA, OrgPath, baseDomain, host)
		if err != nil {
			err = errors.WithMessage(err, "Build peer licenses files error")
			return err
		}
	}

	// 4. check user files
	// 5. make user files
	if len(user) > 0 {
		err = subBuildNewPeerUserInMSP(signCA, tlsCA, OrgPath, baseDomain, user)
		if err != nil {
			err = errors.WithMessage(err, "Build peer licenses files error")
			return err
		}
	}

	return nil
}

func subBuildNewPeerInMSP(signca, tlsca *ca.CA, OrgPath, baseDomain, host string) (err error) {
	// OrgPath /---
	//           |- peers /- host.baseDomain: should not be exist

	url := host + "." + baseDomain
	path := filepath.Join(OrgPath, "peers", url)
	_, err = os.Stat(filepath.Join(path, "msp"))
	if err == nil {
		return errors.Errorf("Peer [%s] is exist", url)
	}

	err = msp.GenerateLocalMSP(path, url, []string{url}, signca, tlsca, msp.PEER, false)
	if err != nil {
		err = errors.WithMessage(err, "Build MSP error")
		return
	}

	err = copyAdminCert(filepath.Join(OrgPath, "users"), filepath.Join(path, "msp", "admincerts"), "Admin@"+baseDomain)
	if err != nil {
		err = errors.WithMessage(err, "Copy Admin user cert error")
		return
	}
	return nil
}

func subBuildNewPeerUserInMSP(signca, tlsca *ca.CA, OrgPath, baseDomain, user string) (err error) {
	uri := user + "@" + baseDomain
	path := filepath.Join(OrgPath, "users", uri)
	_, err = os.Stat(filepath.Join(path, "msp"))
	if err == nil {
		return errors.Errorf("User [%s] is exist", uri)
	}

	err = msp.GenerateLocalMSP(path, uri, []string{}, signca, tlsca, msp.CLIENT, false)
	if err != nil {
		err = errors.WithMessage(err, "build MSP error")
		return
	}

	return nil
}
