package config

// DefaultValues is the default configuration
const DefaultValues = `
[Log]
Level = "debug"
Outputs = ["stdout"]

[Database]
Database = "postgres"
User = "test_user"
Password = "test_password"
Name = "test_db"
Host = "localhost"
Port = "5433"
MaxConns = 20

[Etherman]
L1URL = "http://localhost:8545"
L2URLs = ["http://localhost:8123"]
PrivateKeyPath = "./test/test.keystore"
PrivateKeyPassword = "testonly"

[Synchronizer]
SyncInterval = "2s"
SyncChunkSize = 100
GrpcURL = "localhost:61090"

[BridgeController]
Store = "postgres"
Height = 32

[BridgeServer]
GRPCPort = "9090"
HTTPPort = "8080"
`
