rly chains add --file ./bandchain.json band-laozi-testnet6
rly chains add --file ./buchain.json buchain

# bu-key and band-key already initialize with some balance
rly keys restore band-laozi-testnet6 band-key "ripple ginger valid note naive spray element tube jeans room april real utility join hole cigar double surround already voice risk foot accident panda"
rly keys restore buchain bu-key "genre window palace raccoon youth else wear receive orange heart urge usage image dove viable wrap muffin across trick weird lady cinnamon fiscal boss"

# Add path
rly paths new buchain band-laozi-testnet6 bu-band --dst-port oracle --src-port goldoracle --version bandchain-1

# Link
rly tx link bu-band --src-port goldoracle --dst-port oracle --version bandchain-1
