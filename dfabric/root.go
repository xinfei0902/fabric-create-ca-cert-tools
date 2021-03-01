package dfabric

import (
	"io"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric/common/tools/cryptogen/ca"
	"github.com/pkg/errors"
)

func LoadCAObjectFromFiles(OrgPath string, baseDomain string) (signCA *ca.CA, tlsCA *ca.CA, err error) {
	signCA, err = GetCA(filepath.Join(OrgPath, "ca"), DefaultCASpec("ca."+baseDomain))
	if err != nil {
		err = errors.WithMessage(err, "load ca error")
		return
	}

	tlsCA, err = GetCA(filepath.Join(OrgPath, "tlsca"), DefaultCASpec("tlsca."+baseDomain))
	if err != nil {
		err = errors.WithMessage(err, "load tlsca error")
		return
	}
	return
}

func copyAdminCert(usersDir, adminCertsDir, adminUserName string) error {
	if _, err := os.Stat(filepath.Join(adminCertsDir,
		adminUserName+"-cert.pem")); err == nil {
		return nil
	}
	// delete the contents of admincerts
	err := os.RemoveAll(adminCertsDir)
	if err != nil {
		return err
	}
	// recreate the admincerts directory
	err = os.MkdirAll(adminCertsDir, 0777)
	if err != nil {
		return err
	}
	err = copyFile(filepath.Join(usersDir, adminUserName, "msp", "signcerts",
		adminUserName+"-cert.pem"), filepath.Join(adminCertsDir,
		adminUserName+"-cert.pem"))
	if err != nil {
		return err
	}
	return nil

}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}
