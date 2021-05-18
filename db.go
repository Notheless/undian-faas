package p

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

//NewDBClient to create new connection
func NewDBClient() (*sql.DB, error) {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	ip := os.Getenv("DB_IP")
	rootCert := os.Getenv("DB_CERT_KEY")
	clientCert := os.Getenv("DB_CLIENT_CERT")
	serverName := os.Getenv("DB_SDERVER_NAME")

	rootCertPool := x509.NewCertPool()
	rootCertPool.AppendCertsFromPEM([]byte(rootCert))
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM([]byte(clientCert))

	// rootCertPool := x509.NewCertPool()
	// if ok := rootCertPool.AppendCeratsFromPEM([]byte(pem)); !ok {
	// 	log.Fatal("Failed to append PEM.", user, pass, ip, pem)
	// }
	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs:    rootCertPool,
		ClientCAs:  clientCertPool,
		ServerName: serverName,
	})

	connstring := fmt.Sprintf("%s:%s@%s/default?tls=custom", user, pass, ip)
	db, err := sql.Open("mysql", connstring)
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
