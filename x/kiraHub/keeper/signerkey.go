package keeper

import (
	"errors"
	"strings"
	"time"

	"github.com/KiraCore/sekai/x/kiraHub/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PrefixKeySignerKeys describes the key where to save in KVStore
const PrefixKeySignerKeys = "signer_keys"

// PrefixKeyPubKeyCurator describes the owner of each pubKey
const PrefixKeyPubKeyCurator = "pub_key_curator"

// GetSignerKeys return SignerKeys by a curator
func (k Keeper) GetSignerKeys(ctx sdk.Context, curator sdk.AccAddress) []types.SignerKey {

	var signerKeys []types.SignerKey

	store := ctx.KVStore(k.storeKey)
	curatorStoreID := append([]byte(PrefixKeySignerKeys), []byte(curator)...)

	if store.Has(curatorStoreID) {
		bz := store.Get(curatorStoreID)
		k.cdc.MustUnmarshalBinaryBare(bz, &signerKeys)
	}

	return signerKeys
}

// UpsertSignerKey create signer key and put it into the keeper
func (k Keeper) UpsertSignerKey(ctx sdk.Context,
	pubKey string,
	keyType types.SignerKeyType,
	expiryTime int64,
	Permissions []int64,
	curator sdk.AccAddress) error {

	var newSignerKeys []types.SignerKey
	// TODO: createOrder/createOrderBook should use block time instead of current timestamp from local computer
	lastBlockTimestamp := ctx.BlockHeader().Time.Unix() + time.Hour.Milliseconds()*24*10

	var signerKey = types.NewSignerKey(pubKey, keyType, expiryTime, true, Permissions, curator)

	var signerKeys []types.SignerKey
	// Storage Logic
	store := ctx.KVStore(k.storeKey)
	curatorStoreID := append([]byte(PrefixKeySignerKeys), []byte(curator)...)

	if store.Has(curatorStoreID) {
		bz := store.Get(curatorStoreID)
		k.cdc.MustUnmarshalBinaryBare(bz, &signerKeys)
	}

	pubKeyStoreID := append([]byte(PrefixKeyPubKeyCurator), []byte(pubKey)...)

	for _, sk := range signerKeys {
		if strings.Compare(sk.PubKey, pubKey) == 0 {
			if keyType == sk.KeyType {
				return errors.New("keyType shouldn't be different for same pub key")
			}
			newSignerKeys = append(newSignerKeys, signerKey)
		} else if sk.ExpiryTime > lastBlockTimestamp { // we use lastBlockTimestamp this is only the time that the nodes can agree on
			newSignerKeys = append(newSignerKeys, sk)
		} else { // Delete pubKey curator when it is expired
			if !store.Has(pubKeyStoreID) {
				return errors.New("pubKey to curator mapping is not set properly at the time of expired key cleanup")
			}
			store.Delete(pubKeyStoreID)
		}
	}
	if !store.Has([]byte(pubKey)) { // when pubKey's owner does not exist
		newSignerKeys = append(newSignerKeys, signerKey)
		// Set pubKey curator when pubKey is newly added
		store.Set(pubKeyStoreID, curator)
	} else {
		originCurator := store.Get(pubKeyStoreID)
		if !curator.Equals(sdk.AccAddress(originCurator)) {
			return errors.New("this key is owned by another curator already")
		}
	}

	store.Set(curatorStoreID, k.cdc.MustMarshalBinaryBare(newSignerKeys))
	return nil
}

// TODO: should add test for creating / updating after v0.0.5 release.
// TODO: should add deleteSignerKey after discussion but this should create another directory under transactions folder?