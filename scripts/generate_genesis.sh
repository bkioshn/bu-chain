DIR=`dirname "$0"`

CHAIN_ID="buchain"
DENOM="ububu"

rm -rf ~/.bu-chain

# initial new node
bu-chaind init alice --chain-id $CHAIN_ID
echo "lock nasty suffer dirt dream fine fall deal curtain plate husband sound tower mom crew crawl guard rack snake before fragile course bacon range" \
    | bu-chaind keys add alice --recover --keyring-backend test
echo "smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic" \
    | bu-chaind keys add bob --recover --keyring-backend test


# add accounts to genesis
bu-chaind add-genesis-account alice 10000000000000${DENOM},100000000ngum --keyring-backend test
bu-chaind add-genesis-account bob 10000000000000${DENOM} --keyring-backend test

# register initial validators
bu-chaind gentx alice 100000000${DENOM} \
    --chain-id $CHAIN_ID \
    --keyring-backend test

# collect genesis transactions
bu-chaind collect-gentxs
