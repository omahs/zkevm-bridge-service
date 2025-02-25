package server

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/0xPolygonHermez/zkevm-bridge-service/test/client"
	"github.com/0xPolygonHermez/zkevm-bridge-service/test/operations"
	"github.com/stretchr/testify/require"
)

const (
	grpcPort = "9090"
	restPort = "8080"
)

func init() {
	// Change dir to project root
	// This is important because we have relative paths to files containing test vectors
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestBridgeMock(t *testing.T) {
	_, err := RunMockServer()
	require.NoError(t, err)

	err = operations.WaitGRPCHealthy("0.0.0.0:" + grpcPort)
	require.NoError(t, err)

	url := "http://localhost:" + restPort
	err = operations.WaitRestHealthy(url)
	require.NoError(t, err)

	restClient := client.NewRestClient(url)

	version, err := restClient.GetVersion()
	require.NoError(t, err)
	require.Equal(t, "v1", version)

	offset := uint(0)
	limit := uint(100)
	deposits, totalCount, err := restClient.GetBridges("0xeB17ce701E9D92724AA2ABAdA7E4B28830597Dd9", offset, limit)
	require.NoError(t, err)
	require.Equal(t, len(deposits), 1)
	require.Equal(t, deposits[0].DepositCnt, uint64(4))
	require.Equal(t, totalCount, uint64(1))

	claims, totalCount, err := restClient.GetClaims("0xabCcEd19d7f290B84608feC510bEe872CC8F5112", offset, limit)
	require.NoError(t, err)
	require.Equal(t, len(claims), 1)
	require.Equal(t, totalCount, uint64(1))

	proof, err := restClient.GetMerkleProof(0, 2)
	require.NoError(t, err)
	require.Equal(t, len(proof.MerkleProof), 32)

	deposit, err := restClient.GetBridge(0, 2)
	require.NoError(t, err)
	require.NotEmpty(t, deposit.Metadata)
	require.Equal(t, deposit.ReadyForClaim, true)
	require.Equal(t, deposit.DepositCnt, uint64(2))

	wrappedToken, err := restClient.GetWrappedToken(0, "0x617b3a3528F9cDd6630fd3301B9c8911F7Bf063D")
	require.NoError(t, err)
	require.Equal(t, wrappedToken.WrappedTokenAddr, "0xc2716D3537EcA4B318e60f3d7d6a48714f1F3335")
	require.Equal(t, wrappedToken.Name, "CoinA")
	require.Equal(t, wrappedToken.Symbol, "COA")
	require.Equal(t, wrappedToken.Decimals, uint32(12))
}
