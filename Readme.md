# Calculator

```shell
    docker-compose up --build calc
    time curl  "http://127.0.0.1:8080/api/v1/convert?from=USD&to=USDT&amount=112321321"
```

http://127.0.0.1:8080/swagger/index.html

## Tech
- [Go](https://go.dev)
- [PostgreSQL](https://www.postgresql.org)
- [Docker](https://www.docker.com/)
- [Fiber](https://github.com/gofiber/fiber)
- [Zerolog](https://github.com/rs/zerolog)

## To-Do List
- [😅] XXXXX
- [🐵] XXXXX
- [🐱] XXXXX


## Dev

```shell
    swag init -g internal/app/app.go --output internal/docs
```


## Algorithm


Golang implementation of fastforex.io API data consumer

Pairs are filtered by Postgresql DB.
Only crypto/fiat or fiat/crypto pairs are supported intentionally.


Supported api key by fastforex.io:
```
api_key=be05ce6a0a-f88ddd6bfc-XXXXXX
```


Get available fiat currencies:
```shell
curl --request GET \
     --url 'https://api.fastforex.io/currencies?api_key=be05ce6a0a-f88ddd6bfc-XXXXXX' \
     --header 'accept: application/json'
```


```json
{
  "currencies": {
    "AED": "United Arab Emirates Dirham",
    "AFN": "Afghan Afghani",
    "ALL": "Albanian Lek",
    "AMD": "Armenian Dram",
    "ANG": "Dutch Guilders",
    "AOA": "Angolan Kwanza",
    "ARS": "Argentine Peso",
    "AUD": "Australian Dollar",
    "AWG": "Aruban Florin",
    "AZN": "Azerbaijani Manat",
    "BAM": "Bosnia-Herzegovina Convertible Mark",
    "BBD": "Barbadian Dollar",
    "BDT": "Bangladeshi Taka",
    "BGN": "Bulgarian Lev",
    "BHD": "Bahraini Dinar",
    "BIF": "Burundian Franc",
    "BMD": "Bermudian Dollar",
    "BND": "Bruneian Dollar",
    "BOB": "Bolivian Boliviano",
    "BRL": "Brazilian Real",
    "BSD": "Bahamian Dollar",
    "BTN": "Bhutanese Ngultrum",
    "BWP": "Botswanan Pula",
    "BZD": "Belizean Dollar",
    "CAD": "Canadian Dollar",
    "CDF": "Congolese Franc",
    "CHF": "Swiss Franc",
    "CLF": "Chilean Unit of Account UF",
    "CLP": "Chilean Peso",
    "CNH": "Chinese Yuan Offshore",
    "CNY": "Chinese Yuan",
    "COP": "Colombian Peso",
    "CUP": "Cuban Peso",
    "CVE": "Cape Verdean Escudo",
    "CZK": "Czech Republic Koruna",
    "DJF": "Djiboutian Franc",
    "DKK": "Danish Krone",
    "DOP": "Dominican Peso",
    "DZD": "Algerian Dinar",
    "EGP": "Egyptian Pound",
    "ERN": "Eritrean Nakfa",
    "ETB": "Ethiopian Birr",
    "EUR": "Euro",
    "FJD": "Fijian Dollar",
    "FKP": "Falkland Islands Pound",
    "GBP": "British Pound Sterling",
    "GEL": "Georgian Lari",
    "GHS": "Ghanaian Cedi",
    "GIP": "Gibraltar Pound",
    "GMD": "Gambian Dalasi",
    "GNF": "Guinean Franc",
    "GTQ": "Guatemalan Quetzal",
    "GYD": "Guyanaese Dollar",
    "HKD": "Hong Kong Dollar",
    "HNL": "Honduran Lempira",
    "HRK": "Croatian Kuna",
    "HTG": "Haitian Gourde",
    "HUF": "Hungarian Forint",
    "IDR": "Indonesian Rupiah",
    "ILS": "Israeli New Sheqel",
    "INR": "Indian Rupee",
    "IQD": "Iraqi Dinar",
    "IRR": "Iranian Rial",
    "ISK": "Icelandic Krona",
    "JMD": "Jamaican Dollar",
    "JOD": "Jordanian Dinar",
    "JPY": "Japanese Yen",
    "KES": "Kenyan Shilling",
    "KGS": "Kyrgystani Som",
    "KHR": "Cambodian Riel",
    "KMF": "Comorian Franc",
    "KPW": "North Korean Won",
    "KRW": "South Korean Won",
    "KWD": "Kuwaiti Dinar",
    "KYD": "Caymanian Dollar",
    "KZT": "Kazakhstani Tenge",
    "LAK": "Laotian Kip",
    "LBP": "Lebanese Pound",
    "LKR": "Sri Lankan Rupee",
    "LRD": "Liberian Dollar",
    "LSL": "Basotho Maloti",
    "LYD": "Libyan Dinar",
    "MAD": "Moroccan Dirham",
    "MDL": "Moldovan Leu",
    "MGA": "Malagasy Ariary",
    "MKD": "Macedonian Denar",
    "MMK": "Myanma Kyat",
    "MNT": "Mongolian Tugrik",
    "MOP": "Macanese Pataca",
    "MRU": "Mauritanian Ouguiya",
    "MUR": "Mauritian Rupee",
    "MVR": "Maldivian Rufiyaa",
    "MWK": "Malawian Kwacha",
    "MXN": "Mexican Peso",
    "MYR": "Malaysian Ringgit",
    "MZN": "Mozambican Metical",
    "NAD": "Namibian Dollar",
    "NGN": "Nigerian Naira",
    "NOK": "Norwegian Krone",
    "NPR": "Nepalese Rupee",
    "NZD": "New Zealand Dollar",
    "OMR": "Omani Rial",
    "PAB": "Panamanian Balboa",
    "PEN": "Peruvian Nuevo Sol",
    "PGK": "Papua New Guinean Kina",
    "PHP": "Philippine Peso",
    "PKR": "Pakistani Rupee",
    "PLN": "Polish Zloty",
    "PYG": "Paraguayan Guarani",
    "QAR": "Qatari Rial",
    "RON": "Romanian Leu",
    "RSD": "Serbian Dinar",
    "RUB": "Russian Ruble",
    "RWF": "Rwandan Franc",
    "SAR": "Saudi Arabian Riyal",
    "SCR": "Seychellois Rupee",
    "SDG": "Sudanese Pound",
    "SEK": "Swedish Krona",
    "SGD": "Singapore Dollar",
    "SHP": "Saint Helena Pound",
    "SLL": "Sierra Leonean Leone",
    "SOS": "Somali Shilling",
    "SRD": "Surinamese Dollar",
    "SYP": "Syrian Pound",
    "SZL": "Swazi Emalangeni",
    "THB": "Thai Baht",
    "TJS": "Tajikistani Somoni",
    "TMT": "Turkmenistani Manat",
    "TND": "Tunisian Dinar",
    "TOP": "Tongan Pa'anga",
    "TRY": "Turkish Lira",
    "TTD": "Trinidad and Tobago Dollar",
    "TWD": "Taiwan New Dollar",
    "TZS": "Tanzanian Shilling",
    "UAH": "Ukrainian Hryvnia",
    "UGX": "Ugandan Shilling",
    "USD": "United States Dollar",
    "UYU": "Uruguayan Peso",
    "UZS": "Uzbekistan Som",
    "VND": "Vietnamese Dong",
    "VUV": "Ni-Vanuatu Vatu",
    "WST": "Samoan Tala",
    "XAF": "CFA Franc BEAC",
    "XCD": "East Caribbean Dollar",
    "XDR": "Special Drawing Rights",
    "XOF": "CFA Franc BCEAO",
    "XPF": "CFP Franc",
    "YER": "Yemeni Rial",
    "ZAR": "South African Rand",
    "ZMW": "Zambian Kwacha"
  },
  "ms": 1
}

```

Get available crypto currencies:
```shell
curl --request GET \
     --url 'https://api.fastforex.io/crypto/currencies?api_key=be05ce6a0a-f88ddd6bfc-XXXXXX' \
     --header 'accept: application/json'

```
310 crypto currencies:

```json
    {
  "currencies": {
    "1INCH": "1inch Network",
    "AAVE": "Aave",
    "ACA": "Acala Token",
    "ACH": "Alchemy Pay",
    "ACM": "AC Milan Fan Token",
    "ADA": "Cardano",
    "ADX": "Ambire AdEx",
    "AERGO": "Aergo",
    "AGIX": "SingularityNET",
    "AGLD": "Adventure Gold",
    "AIOZ": "AIOZ Network",
    "AKRO": "Akropolis",
    "ALCX": "Alchemix",
    "ALGO": "Algorand",
    "ALICE": "MyNeighborAlice",
    "ALPHA": "Alpha Venture DAO",
    "ALPINE": "Alpine F1 Team Fan Token",
    "AMB": "AirDAO",
    "AMP": "Amp",
    "ANKR": "Ankr",
    "ANT": "Aragon",
    "APE": "ApeCoin",
    "API3": "API3",
    "APT": "Aptos",
    "AR": "Arweave",
    "ARPA": "ARPA",
    "ASTR": "Astar",
    "ATA": "Automata Network",
    "ATOM": "Cosmos",
    "AUDIO": "Audius",
    "AURORA": "Aurora",
    "AVA": "Travala.com",
    "AVAX": "Avalanche",
    "AXS": "Axie Infinity",
    "BADGER": "Badger DAO",
    "BAL": "Balancer",
    "BAND": "Band Protocol",
    "BAT": "Basic Attention Token",
    "BCH": "Bitcoin Cash",
    "BEL": "Bella Protocol",
    "BICO": "Biconomy",
    "BIFI": "Beefy Finance",
    "BLZ": "Bluzelle",
    "BNB": "BNB",
    "BNT": "Bancor",
    "BOBA": "Boba Network",
    "BOND": "BarnBridge",
    "BOSON": "Boson Protocol",
    "BSW": "Biswap",
    "BTC": "Bitcoin",
    "BTT": "BitTorrent-New",
    "BUSD": "Binance USD",
    "C98": "Coin98",
    "CAKE": "PancakeSwap",
    "CELO": "Celo",
    "CELR": "Celer Network",
    "CFX": "Conflux",
    "CHR": "Chromia",
    "CHZ": "Chiliz",
    "CITY": "Manchester City Fan Token",
    "CKB": "Nervos Network",
    "CLV": "CLV",
    "COMP": "Compound",
    "COTI": "COTI",
    "CQT": "Covalent",
    "CREAM": "Cream Finance",
    "CRO": "Cronos",
    "CRV": "Curve DAO Token",
    "CSPR": "Casper",
    "CTC": "Creditcoin",
    "CTSI": "Cartesi",
    "CUDOS": "CUDOS",
    "CULT": "Cult DAO",
    "CVX": "Convex Finance",
    "DAI": "Dai",
    "DAR": "Mines of Dalarnia",
    "DASH": "Dash",
    "DATA": "Streamr",
    "DCR": "Decred",
    "DEGO": "Dego Finance",
    "DENT": "Dent",
    "DERC": "DeRace",
    "DEXE": "DeXe",
    "DFI": "DeFiChain",
    "DGB": "DigiByte",
    "DIA": "DIA",
    "DODO": "DODO",
    "DOGE": "Dogecoin",
    "DOT": "Polkadot",
    "DUSK": "Dusk Network",
    "DYDX": "dYdX",
    "EGLD": "Elrond",
    "ELF": "aelf",
    "ELON": "Dogelon Mars",
    "ENJ": "Enjin Coin",
    "ENS": "Ethereum Name Service",
    "EOS": "EOS",
    "ETC": "Ethereum Classic",
    "ETH": "Ethereum",
    "ETHW": "EthereumPoW",
    "EVER": "Everscale",
    "FARM": "Harvest Finance",
    "FET": "Fetch.ai",
    "FIDA": "Bonfida",
    "FIL": "Filecoin",
    "FIS": "StaFi",
    "FLOW": "Flow",
    "FLUX": "Flux",
    "FORT": "Forta",
    "FORTH": "Ampleforth Governance Token",
    "FRONT": "Frontier",
    "FTM": "Fantom",
    "FTT": "FTX Token",
    "FXS": "Frax Share",
    "GAL": "Galatasaray Fan Token",
    "GALA": "Gala",
    "GAS": "Gas",
    "GHST": "Aavegotchi",
    "GLM": "Golem",
    "GLMR": "Moonbeam",
    "GMT": "GMT Token",
    "GMX": "GMX",
    "GNO": "Gnosis",
    "GODS": "Gods Unchained",
    "GRT": "The Graph",
    "GST": "Green Satoshi Token (SOL)",
    "GTC": "Gitcoin",
    "HARD": "Kava Lend",
    "HBAR": "Hedera",
    "HERO": "Metahero",
    "HIGH": "Highstreet",
    "HNT": "Helium",
    "HOT": "Holo",
    "ICP": "Internet Computer",
    "ICX": "ICON",
    "ID": "Everest",
    "IDEX": "IDEX",
    "ILV": "Illuvium",
    "IMX": "Immutable X",
    "INJ": "Injective",
    "IOST": "IOST",
    "IOTX": "IoTeX",
    "JASMY": "JasmyCoin",
    "JOE": "JOE",
    "JST": "JUST",
    "JUV": "Juventus Fan Token",
    "KAS": "Kaspa",
    "KAVA": "Kava",
    "KDA": "Kadena",
    "KLAY": "Klaytn",
    "KNC": "Kyber Network Crystal v2",
    "KOK": "KOK",
    "KRL": "Kryll",
    "KSM": "Kusama",
    "LDO": "Lido DAO",
    "LEVER": "LeverFi",
    "LINA": "Linear Finance",
    "LINK": "Chainlink",
    "LIT": "Litentry",
    "LOKA": "League of Kingdoms Arena",
    "LOOKS": "LooksRare",
    "LPT": "Livepeer",
    "LQTY": "Liquity",
    "LRC": "Loopring",
    "LSK": "Lisk",
    "LTC": "Litecoin",
    "LTO": "LTO Network",
    "LUNA": "Terra",
    "LUNC": "Terra Classic",
    "MAGIC": "MAGIC",
    "MANA": "Decentraland",
    "MASK": "Mask Network",
    "MATIC": "Polygon",
    "MBL": "MovieBloc",
    "MDT": "Measurable Data Token",
    "METIS": "MetisDAO",
    "MINA": "Mina",
    "MKR": "Maker",
    "MLN": "Enzyme",
    "MOVR": "Moonriver",
    "MTL": "Metal DAO",
    "MV": "GensoKishi Metaverse",
    "MXC": "MXC",
    "NEAR": "NEAR Protocol",
    "NEO": "Neo",
    "NEXO": "Nexo",
    "NFT": "APENFT",
    "NKN": "NKN",
    "NMR": "Numeraire",
    "OCEAN": "Ocean Protocol",
    "OGN": "Origin Protocol",
    "OM": "MANTRA",
    "OMG": "OMG Network",
    "ONE": "BigONE Token",
    "ONT": "Ontology",
    "OP": "Optimism",
    "ORBS": "Orbs",
    "ORCA": "Orca",
    "ORN": "Orion Protocol",
    "OSMO": "Osmosis",
    "OXT": "Orchid",
    "PAXG": "PAX Gold",
    "PENDLE": "Pendle",
    "PEOPLE": "ConstitutionDAO",
    "PERP": "Perpetual Protocol",
    "PHA": "Phala Network",
    "PLA": "PlayDapp",
    "POLS": "Polkastarter",
    "POND": "Marlin",
    "POWR": "Powerledger",
    "PROM": "Prom",
    "PSG": "Paris Saint-Germain Fan Token",
    "PSTAKE": "pSTAKE Finance",
    "PUNDIX": "Pundi X (New)",
    "PYR": "Vulcan Forged PYR",
    "QI": "BENQI",
    "QKC": "QuarkChain",
    "QNT": "Quant",
    "QTUM": "Qtum",
    "QUICK": "QuickSwap",
    "RACA": "RadioCaca",
    "RAD": "Radicle",
    "RARE": "SuperRare",
    "RARI": "Rarible",
    "REEF": "Reef",
    "REN": "Ren",
    "REQ": "Request",
    "RLC": "iExec RLC",
    "RNDR": "Render Token",
    "ROSE": "Oasis Network",
    "RPL": "Rocket Pool",
    "RSR": "Reserve Rights",
    "RUNE": "THORChain",
    "RVN": "Ravencoin",
    "SAND": "The Sandbox",
    "SC": "Siacoin",
    "SCRT": "Secret",
    "SFP": "SafePal",
    "SFUND": "Seedify.fund",
    "SHIB": "Shiba Inu",
    "SHILL": "SHILL Token",
    "SIDUS": "SIDUS",
    "SKL": "SKALE Network",
    "SLP": "Smooth Love Potion",
    "SNX": "Synthetix",
    "SOL": "Solana",
    "SPELL": "Spell Token",
    "SSV": "ssv.network",
    "STG": "Stargate Finance",
    "STORJ": "Storj",
    "STRAX": "Stratis",
    "STX": "Stacks",
    "SUN": "Sun (New)",
    "SUPER": "SuperFarm",
    "SUSHI": "SushiSwap",
    "SWEAT": "Sweat Economy",
    "SXP": "SXP",
    "SYN": "Synapse",
    "SYS": "Syscoin",
    "T": "Threshold",
    "TFUEL": "Theta Fuel",
    "THETA": "Theta Network",
    "TLM": "Alien Worlds",
    "TON": "Tokamak Network",
    "TRAC": "OriginTrail",
    "TRB": "Tellor",
    "TRU": "TrueFi",
    "TRVL": "TRVL",
    "TRX": "TRON",
    "TUSD": "TrueUSD",
    "TVK": "Virtua",
    "TWT": "Trust Wallet Token",
    "UMA": "UMA",
    "UNFI": "Unifi Protocol DAO",
    "UNI": "UNICORN Token",
    "USDC": "USD Coin",
    "USDD": "USDD",
    "USDP": "Pax Dollar",
    "USDT": "Tether",
    "USTC": "TerraClassicUSD",
    "VEGA": "Vega Protocol",
    "VELO": "Velo",
    "VET": "VeChain",
    "VOXEL": "Voxies",
    "VRA": "Verasity",
    "WAVES": "Waves",
    "WAXP": "WAX",
    "WBTC": "Wrapped Bitcoin",
    "WEMIX": "WEMIX",
    "WIN": "WINkLink",
    "WOO": "WOO Network",
    "WRX": "WazirX",
    "XCAD": "XCAD Network",
    "XDC": "XDC Network",
    "XEC": "eCash",
    "XEM": "NEM",
    "XLM": "Stellar",
    "XMR": "Monero",
    "XNO": "Xeno Token",
    "XRP": "XRP",
    "XTZ": "Tezos",
    "XWG": "X World Games",
    "XYM": "Symbol",
    "XYO": "XYO",
    "YFI": "yearn.finance",
    "YGG": "Yield Guild Games",
    "ZEC": "Zcash",
    "ZEN": "Horizen",
    "ZIL": "Zilliqa",
    "ZRX": "0x"
  },
  "ms": 4
}
```

Fetch a list of 400+ supported cryptocurrency pairs

```shell
curl --request GET \
     --url 'https://api.fastforex.io/crypto/pairs?api_key=be05ce6a0a-f88ddd6bfc-XXXXXX' \
     --header 'accept: application/json'
```

630 cryptocurrency pairs

```
{
  "pairs": {
    "1INCH/USD": {
      "base": "1INCH",
      "quote": "USD"
    },
    "1INCH/USDT": {
      "base": "1INCH",
      "quote": "USDT"
    },
    "AAVE/BTC": {
      "base": "AAVE",
      "quote": "BTC"
    },
    "AAVE/ETH": {
      "base": "AAVE",
      "quote": "ETH"
    },
    "AAVE/USD": {
      "base": "AAVE",
      "quote": "USD"
    },
    "AAVE/USDT": {
      "base": "AAVE",
      "quote": "USDT"
    },
    "ACA/USDT": {
      "base": "ACA",
      "quote": "USDT"
    },
    "ACH/USD": {
      "base": "ACH",
      "quote": "USD"
    },
    "ACH/USDT": {
      "base": "ACH",
      "quote": "USDT"
    },
    "ACM/USDT": {
      "base": "ACM",
      "quote": "USDT"
    },
    "ACS/USDT": {
      "base": "ACS",
      "quote": "USDT"
    },
    "ADA/BTC": {
      "base": "ADA",
      "quote": "BTC"
    },
    "ADA/ETH": {
      "base": "ADA",
      "quote": "ETH"
    },
    "ADA/EUR": {
      "base": "ADA",
      "quote": "EUR"
    },
    "ADA/GBP": {
      "base": "ADA",
      "quote": "GBP"
    },
    "ADA/USD": {
      "base": "ADA",
      "quote": "USD"
    },
    "ADA/USDC": {
      "base": "ADA",
      "quote": "USDC"
    },
    "ADA/USDT": {
      "base": "ADA",
      "quote": "USDT"
    },
    "ADX/USDT": {
      "base": "ADX",
      "quote": "USDT"
    },
    "AERGO/USDT": {
      "base": "AERGO",
      "quote": "USDT"
    },
    "AGIX/USDT": {
      "base": "AGIX",
      "quote": "USDT"
    },
    "AGLD/USD": {
      "base": "AGLD",
      "quote": "USD"
    },
    "AGLD/USDT": {
      "base": "AGLD",
      "quote": "USDT"
    },
    "AIOZ/USDT": {
      "base": "AIOZ",
      "quote": "USDT"
    },
    "AKRO/USDT": {
      "base": "AKRO",
      "quote": "USDT"
    },
    "ALCX/USD": {
      "base": "ALCX",
      "quote": "USD"
    },
    "ALGO/BTC": {
      "base": "ALGO",
      "quote": "BTC"
    },
    "ALGO/ETH": {
      "base": "ALGO",
      "quote": "ETH"
    },
    "ALGO/USD": {
      "base": "ALGO",
      "quote": "USD"
    },
    "ALGO/USDT": {
      "base": "ALGO",
      "quote": "USDT"
    },
    "ALICE/USD": {
      "base": "ALICE",
      "quote": "USD"
    },
    "ALICE/USDT": {
      "base": "ALICE",
      "quote": "USDT"
    },
    "ALPHA/USDT": {
      "base": "ALPHA",
      "quote": "USDT"
    },
    "ALPINE/USDT": {
      "base": "ALPINE",
      "quote": "USDT"
    },
    "AMB/USDT": {
      "base": "AMB",
      "quote": "USDT"
    },
    "AMP/USDT": {
      "base": "AMP",
      "quote": "USDT"
    },
    "ANKR/BTC": {
      "base": "ANKR",
      "quote": "BTC"
    },
    "ANKR/USD": {
      "base": "ANKR",
      "quote": "USD"
    },
    "ANKR/USDT": {
      "base": "ANKR",
      "quote": "USDT"
    },
    "ANT/BTC": {
      "base": "ANT",
      "quote": "BTC"
    },
    "ANT/USDT": {
      "base": "ANT",
      "quote": "USDT"
    },
    "APE/EUR": {
      "base": "APE",
      "quote": "EUR"
    },
    "APE/USD": {
      "base": "APE",
      "quote": "USD"
    },
    "APE/USDT": {
      "base": "APE",
      "quote": "USDT"
    },
    "API3/USD": {
      "base": "API3",
      "quote": "USD"
    },
    "API3/USDT": {
      "base": "API3",
      "quote": "USDT"
    },
    "APT/USD": {
      "base": "APT",
      "quote": "USD"
    },
    "APT/USDT": {
      "base": "APT",
      "quote": "USDT"
    },
    "AR/USDT": {
      "base": "AR",
      "quote": "USDT"
    },
    "ARB/USD": {
      "base": "ARB",
      "quote": "USD"
    },
    "ARB/USDT": {
      "base": "ARB",
      "quote": "USDT"
    },
    "ARKM/USDT": {
      "base": "ARKM",
      "quote": "USDT"
    },
    "ARPA/USD": {
      "base": "ARPA",
      "quote": "USD"
    },
    "ARPA/USDT": {
      "base": "ARPA",
      "quote": "USDT"
    },
    "ASTR/BTC": {
      "base": "ASTR",
      "quote": "BTC"
    },
    "ASTR/USDT": {
      "base": "ASTR",
      "quote": "USDT"
    },
    "ATA/USDT": {
      "base": "ATA",
      "quote": "USDT"
    },
    "ATOM/BTC": {
      "base": "ATOM",
      "quote": "BTC"
    },
    "ATOM/ETH": {
      "base": "ATOM",
      "quote": "ETH"
    },
    "ATOM/EUR": {
      "base": "ATOM",
      "quote": "EUR"
    },
    "ATOM/USD": {
      "base": "ATOM",
      "quote": "USD"
    },
    "ATOM/USDT": {
      "base": "ATOM",
      "quote": "USDT"
    },
    "AUDIO/USD": {
      "base": "AUDIO",
      "quote": "USD"
    },
    "AUDIO/USDT": {
      "base": "AUDIO",
      "quote": "USDT"
    },
    "AURORA/USDT": {
      "base": "AURORA",
      "quote": "USDT"
    },
    "AVA/USDT": {
      "base": "AVA",
      "quote": "USDT"
    },
    "AVAX/BTC": {
      "base": "AVAX",
      "quote": "BTC"
    },
    "AVAX/EUR": {
      "base": "AVAX",
      "quote": "EUR"
    },
    "AVAX/USD": {
      "base": "AVAX",
      "quote": "USD"
    },
    "AVAX/USDC": {
      "base": "AVAX",
      "quote": "USDC"
    },
    "AVAX/USDT": {
      "base": "AVAX",
      "quote": "USDT"
    },
    "AXL/USDT": {
      "base": "AXL",
      "quote": "USDT"
    },
    "AXS/USD": {
      "base": "AXS",
      "quote": "USD"
    },
    "AXS/USDT": {
      "base": "AXS",
      "quote": "USDT"
    },
    "BABYDOGE/USDT": {
      "base": "BABYDOGE",
      "quote": "USDT"
    },
    "BADGER/USD": {
      "base": "BADGER",
      "quote": "USD"
    },
    "BADGER/USDT": {
      "base": "BADGER",
      "quote": "USDT"
    },
    "BAL/BTC": {
      "base": "BAL",
      "quote": "BTC"
    },
    "BAL/USD": {
      "base": "BAL",
      "quote": "USD"
    },
    "BAL/USDT": {
      "base": "BAL",
      "quote": "USDT"
    },
    "BAND/USD": {
      "base": "BAND",
      "quote": "USD"
    },
    "BAND/USDT": {
      "base": "BAND",
      "quote": "USDT"
    },
    "BAT/BTC": {
      "base": "BAT",
      "quote": "BTC"
    },
    "BAT/ETH": {
      "base": "BAT",
      "quote": "ETH"
    },
    "BAT/USD": {
      "base": "BAT",
      "quote": "USD"
    },
    "BAT/USDT": {
      "base": "BAT",
      "quote": "USDT"
    },
    "BCH/BTC": {
      "base": "BCH",
      "quote": "BTC"
    },
    "BCH/EUR": {
      "base": "BCH",
      "quote": "EUR"
    },
    "BCH/USD": {
      "base": "BCH",
      "quote": "USD"
    },
    "BCH/USDT": {
      "base": "BCH",
      "quote": "USDT"
    },
    "BEL/USDT": {
      "base": "BEL",
      "quote": "USDT"
    },
    "BICO/USD": {
      "base": "BICO",
      "quote": "USD"
    },
    "BICO/USDT": {
      "base": "BICO",
      "quote": "USDT"
    },
    "BIFI/USDT": {
      "base": "BIFI",
      "quote": "USDT"
    },
    "BLUR/USD": {
      "base": "BLUR",
      "quote": "USD"
    },
    "BLUR/USDT": {
      "base": "BLUR",
      "quote": "USDT"
    },
    "BLZ/USD": {
      "base": "BLZ",
      "quote": "USD"
    },
    "BLZ/USDT": {
      "base": "BLZ",
      "quote": "USDT"
    },
    "BNB/BTC": {
      "base": "BNB",
      "quote": "BTC"
    },
    "BNB/USDC": {
      "base": "BNB",
      "quote": "USDC"
    },
    "BNB/USDT": {
      "base": "BNB",
      "quote": "USDT"
    },
    "BNT/USD": {
      "base": "BNT",
      "quote": "USD"
    },
    "BNT/USDT": {
      "base": "BNT",
      "quote": "USDT"
    },
    "BOBA/USD": {
      "base": "BOBA",
      "quote": "USD"
    },
    "BOBA/USDT": {
      "base": "BOBA",
      "quote": "USDT"
    },
    "BOND/USDT": {
      "base": "BOND",
      "quote": "USDT"
    },
    "BOSON/USDT": {
      "base": "BOSON",
      "quote": "USDT"
    },
    "BSW/USDT": {
      "base": "BSW",
      "quote": "USDT"
    },
    "BTC/DAI": {
      "base": "BTC",
      "quote": "DAI"
    },
    "BTC/EUR": {
      "base": "BTC",
      "quote": "EUR"
    },
    "BTC/GBP": {
      "base": "BTC",
      "quote": "GBP"
    },
    "BTC/USD": {
      "base": "BTC",
      "quote": "USD"
    },
    "BTC/USDC": {
      "base": "BTC",
      "quote": "USDC"
    },
    "BTC/USDT": {
      "base": "BTC",
      "quote": "USDT"
    },
    "BTC3L/USDT": {
      "base": "BTC3L",
      "quote": "USDT"
    },
    "BTC3S/USDT": {
      "base": "BTC3S",
      "quote": "USDT"
    },
    "BTT/USDT": {
      "base": "BTT",
      "quote": "USDT"
    },
    "BUSD/USDT": {
      "base": "BUSD",
      "quote": "USDT"
    },
    "C98/USD": {
      "base": "C98",
      "quote": "USD"
    },
    "C98/USDT": {
      "base": "C98",
      "quote": "USDT"
    },
    "CAKE/USDT": {
      "base": "CAKE",
      "quote": "USDT"
    },
    "CELO/USDT": {
      "base": "CELO",
      "quote": "USDT"
    },
    "CELR/USD": {
      "base": "CELR",
      "quote": "USD"
    },
    "CELR/USDT": {
      "base": "CELR",
      "quote": "USDT"
    },
    "CFX/USDT": {
      "base": "CFX",
      "quote": "USDT"
    },
    "CGPT/USDT": {
      "base": "CGPT",
      "quote": "USDT"
    },
    "CHR/USDT": {
      "base": "CHR",
      "quote": "USDT"
    },
    "CHZ/EUR": {
      "base": "CHZ",
      "quote": "EUR"
    },
    "CHZ/USD": {
      "base": "CHZ",
      "quote": "USD"
    },
    "CHZ/USDT": {
      "base": "CHZ",
      "quote": "USDT"
    },
    "CITY/USDT": {
      "base": "CITY",
      "quote": "USDT"
    },
    "CKB/USDT": {
      "base": "CKB",
      "quote": "USDT"
    },
    "CLV/USDT": {
      "base": "CLV",
      "quote": "USDT"
    },
    "CMP/USDT": {
      "base": "CMP",
      "quote": "USDT"
    },
    "COMBO/USDT": {
      "base": "COMBO",
      "quote": "USDT"
    },
    "COMP/BTC": {
      "base": "COMP",
      "quote": "BTC"
    },
    "COMP/USD": {
      "base": "COMP",
      "quote": "USD"
    },
    "COMP/USDT": {
      "base": "COMP",
      "quote": "USDT"
    },
    "COTI/USD": {
      "base": "COTI",
      "quote": "USD"
    },
    "COTI/USDT": {
      "base": "COTI",
      "quote": "USDT"
    },
    "CQT/USDT": {
      "base": "CQT",
      "quote": "USDT"
    },
    "CREAM/USDT": {
      "base": "CREAM",
      "quote": "USDT"
    },
    "CRO/USDT": {
      "base": "CRO",
      "quote": "USDT"
    },
    "CRV/BTC": {
      "base": "CRV",
      "quote": "BTC"
    },
    "CRV/ETH": {
      "base": "CRV",
      "quote": "ETH"
    },
    "CRV/USD": {
      "base": "CRV",
      "quote": "USD"
    },
    "CRV/USDT": {
      "base": "CRV",
      "quote": "USDT"
    },
    "CSPR/USDT": {
      "base": "CSPR",
      "quote": "USDT"
    },
    "CTC/USDT": {
      "base": "CTC",
      "quote": "USDT"
    },
    "CTSI/USD": {
      "base": "CTSI",
      "quote": "USD"
    },
    "CTSI/USDT": {
      "base": "CTSI",
      "quote": "USDT"
    },
    "CUDOS/USDT": {
      "base": "CUDOS",
      "quote": "USDT"
    },
    "CULT/USDT": {
      "base": "CULT",
      "quote": "USDT"
    },
    "CVX/USD": {
      "base": "CVX",
      "quote": "USD"
    },
    "CVX/USDT": {
      "base": "CVX",
      "quote": "USDT"
    },
    "CWAR/USDT": {
      "base": "CWAR",
      "quote": "USDT"
    },
    "CYBER/USDT": {
      "base": "CYBER",
      "quote": "USDT"
    },
    "DAI/USD": {
      "base": "DAI",
      "quote": "USD"
    },
    "DAI/USDT": {
      "base": "DAI",
      "quote": "USDT"
    },
    "DAR/USDT": {
      "base": "DAR",
      "quote": "USDT"
    },
    "DASH/BTC": {
      "base": "DASH",
      "quote": "BTC"
    },
    "DASH/USDT": {
      "base": "DASH",
      "quote": "USDT"
    },
    "DATA/USDT": {
      "base": "DATA",
      "quote": "USDT"
    },
    "DCR/BTC": {
      "base": "DCR",
      "quote": "BTC"
    },
    "DCR/USDT": {
      "base": "DCR",
      "quote": "USDT"
    },
    "DEGO/USDT": {
      "base": "DEGO",
      "quote": "USDT"
    },
    "DENT/ETH": {
      "base": "DENT",
      "quote": "ETH"
    },
    "DERC/USDT": {
      "base": "DERC",
      "quote": "USDT"
    },
    "DEXE/ETH": {
      "base": "DEXE",
      "quote": "ETH"
    },
    "DEXE/USDT": {
      "base": "DEXE",
      "quote": "USDT"
    },
    "DFI/USDT": {
      "base": "DFI",
      "quote": "USDT"
    },
    "DGB/USDT": {
      "base": "DGB",
      "quote": "USDT"
    },
    "DIA/USDT": {
      "base": "DIA",
      "quote": "USDT"
    },
    "DODO/USDT": {
      "base": "DODO",
      "quote": "USDT"
    },
    "DOGE/BTC": {
      "base": "DOGE",
      "quote": "BTC"
    },
    "DOGE/EUR": {
      "base": "DOGE",
      "quote": "EUR"
    },
    "DOGE/USD": {
      "base": "DOGE",
      "quote": "USD"
    },
    "DOGE/USDC": {
      "base": "DOGE",
      "quote": "USDC"
    },
    "DOGE/USDT": {
      "base": "DOGE",
      "quote": "USDT"
    },
    "DOT/BTC": {
      "base": "DOT",
      "quote": "BTC"
    },
    "DOT/EUR": {
      "base": "DOT",
      "quote": "EUR"
    },
    "DOT/USD": {
      "base": "DOT",
      "quote": "USD"
    },
    "DOT/USDT": {
      "base": "DOT",
      "quote": "USDT"
    },
    "DOT3L/USDT": {
      "base": "DOT3L",
      "quote": "USDT"
    },
    "DOT3S/USDT": {
      "base": "DOT3S",
      "quote": "USDT"
    },
    "DUSK/USDT": {
      "base": "DUSK",
      "quote": "USDT"
    },
    "DYDX/USDT": {
      "base": "DYDX",
      "quote": "USDT"
    },
    "ECOX/USDT": {
      "base": "ECOX",
      "quote": "USDT"
    },
    "EDU/USDT": {
      "base": "EDU",
      "quote": "USDT"
    },
    "EGLD/USD": {
      "base": "EGLD",
      "quote": "USD"
    },
    "EGLD/USDT": {
      "base": "EGLD",
      "quote": "USDT"
    },
    "ELF/ETH": {
      "base": "ELF",
      "quote": "ETH"
    },
    "ELON/USDT": {
      "base": "ELON",
      "quote": "USDT"
    },
    "ENJ/BTC": {
      "base": "ENJ",
      "quote": "BTC"
    },
    "ENJ/ETH": {
      "base": "ENJ",
      "quote": "ETH"
    },
    "ENJ/USD": {
      "base": "ENJ",
      "quote": "USD"
    },
    "ENJ/USDT": {
      "base": "ENJ",
      "quote": "USDT"
    },
    "ENS/USD": {
      "base": "ENS",
      "quote": "USD"
    },
    "ENS/USDT": {
      "base": "ENS",
      "quote": "USDT"
    },
    "EOS/BTC": {
      "base": "EOS",
      "quote": "BTC"
    },
    "EOS/ETH": {
      "base": "EOS",
      "quote": "ETH"
    },
    "EOS/EUR": {
      "base": "EOS",
      "quote": "EUR"
    },
    "EOS/USD": {
      "base": "EOS",
      "quote": "USD"
    },
    "EOS/USDT": {
      "base": "EOS",
      "quote": "USDT"
    },
    "EPX/USDT": {
      "base": "EPX",
      "quote": "USDT"
    },
    "ETC/BTC": {
      "base": "ETC",
      "quote": "BTC"
    },
    "ETC/ETH": {
      "base": "ETC",
      "quote": "ETH"
    },
    "ETC/EUR": {
      "base": "ETC",
      "quote": "EUR"
    },
    "ETC/USD": {
      "base": "ETC",
      "quote": "USD"
    },
    "ETC/USDT": {
      "base": "ETC",
      "quote": "USDT"
    },
    "ETH/BTC": {
      "base": "ETH",
      "quote": "BTC"
    },
    "ETH/DAI": {
      "base": "ETH",
      "quote": "DAI"
    },
    "ETH/EUR": {
      "base": "ETH",
      "quote": "EUR"
    },
    "ETH/GBP": {
      "base": "ETH",
      "quote": "GBP"
    },
    "ETH/USD": {
      "base": "ETH",
      "quote": "USD"
    },
    "ETH/USDC": {
      "base": "ETH",
      "quote": "USDC"
    },
    "ETH/USDT": {
      "base": "ETH",
      "quote": "USDT"
    },
    "ETH3L/USDT": {
      "base": "ETH3L",
      "quote": "USDT"
    },
    "ETH3S/USDT": {
      "base": "ETH3S",
      "quote": "USDT"
    },
    "ETHW/USDT": {
      "base": "ETHW",
      "quote": "USDT"
    },
    "EVER/USDT": {
      "base": "EVER",
      "quote": "USDT"
    },
    "FARM/USD": {
      "base": "FARM",
      "quote": "USD"
    },
    "FET/USD": {
      "base": "FET",
      "quote": "USD"
    },
    "FET/USDT": {
      "base": "FET",
      "quote": "USDT"
    },
    "FIDA/USDT": {
      "base": "FIDA",
      "quote": "USDT"
    },
    "FIL/BTC": {
      "base": "FIL",
      "quote": "BTC"
    },
    "FIL/ETH": {
      "base": "FIL",
      "quote": "ETH"
    },
    "FIL/USD": {
      "base": "FIL",
      "quote": "USD"
    },
    "FIL/USDT": {
      "base": "FIL",
      "quote": "USDT"
    },
    "FIS/USD": {
      "base": "FIS",
      "quote": "USD"
    },
    "FITFI/USDT": {
      "base": "FITFI",
      "quote": "USDT"
    },
    "FLOKI/USDT": {
      "base": "FLOKI",
      "quote": "USDT"
    },
    "FLOW/BTC": {
      "base": "FLOW",
      "quote": "BTC"
    },
    "FLOW/USD": {
      "base": "FLOW",
      "quote": "USD"
    },
    "FLOW/USDT": {
      "base": "FLOW",
      "quote": "USDT"
    },
    "FLR/USD": {
      "base": "FLR",
      "quote": "USD"
    },
    "FLR/USDT": {
      "base": "FLR",
      "quote": "USDT"
    },
    "FLUX/USDT": {
      "base": "FLUX",
      "quote": "USDT"
    },
    "FORT/USDT": {
      "base": "FORT",
      "quote": "USDT"
    },
    "FORTH/USD": {
      "base": "FORTH",
      "quote": "USD"
    },
    "FORTH/USDT": {
      "base": "FORTH",
      "quote": "USDT"
    },
    "FRONT/USDT": {
      "base": "FRONT",
      "quote": "USDT"
    },
    "FTM/ETH": {
      "base": "FTM",
      "quote": "ETH"
    },
    "FTM/USDT": {
      "base": "FTM",
      "quote": "USDT"
    },
    "FTT/USDT": {
      "base": "FTT",
      "quote": "USDT"
    },
    "FXS/USDT": {
      "base": "FXS",
      "quote": "USDT"
    },
    "GAL/USD": {
      "base": "GAL",
      "quote": "USD"
    },
    "GAL/USDT": {
      "base": "GAL",
      "quote": "USDT"
    },
    "GALA/USDT": {
      "base": "GALA",
      "quote": "USDT"
    },
    "GAS/BTC": {
      "base": "GAS",
      "quote": "BTC"
    },
    "GAS/USDT": {
      "base": "GAS",
      "quote": "USDT"
    },
    "GFT/USDT": {
      "base": "GFT",
      "quote": "USDT"
    },
    "GHST/USD": {
      "base": "GHST",
      "quote": "USD"
    },
    "GLM/USDT": {
      "base": "GLM",
      "quote": "USDT"
    },
    "GLMR/USDT": {
      "base": "GLMR",
      "quote": "USDT"
    },
    "GMT/USD": {
      "base": "GMT",
      "quote": "USD"
    },
    "GMT/USDT": {
      "base": "GMT",
      "quote": "USDT"
    },
    "GMX/USDT": {
      "base": "GMX",
      "quote": "USDT"
    },
    "GNO/USD": {
      "base": "GNO",
      "quote": "USD"
    },
    "GNS/USDT": {
      "base": "GNS",
      "quote": "USDT"
    },
    "GODS/USDT": {
      "base": "GODS",
      "quote": "USDT"
    },
    "GRT/BTC": {
      "base": "GRT",
      "quote": "BTC"
    },
    "GRT/EUR": {
      "base": "GRT",
      "quote": "EUR"
    },
    "GRT/USD": {
      "base": "GRT",
      "quote": "USD"
    },
    "GRT/USDT": {
      "base": "GRT",
      "quote": "USDT"
    },
    "GST/USDT": {
      "base": "GST",
      "quote": "USDT"
    },
    "GTC/BTC": {
      "base": "GTC",
      "quote": "BTC"
    },
    "GTC/USD": {
      "base": "GTC",
      "quote": "USD"
    },
    "GTC/USDT": {
      "base": "GTC",
      "quote": "USDT"
    },
    "HARD/USDT": {
      "base": "HARD",
      "quote": "USDT"
    },
    "HBAR/BTC": {
      "base": "HBAR",
      "quote": "BTC"
    },
    "HBAR/USDT": {
      "base": "HBAR",
      "quote": "USDT"
    },
    "HERO/USDT": {
      "base": "HERO",
      "quote": "USDT"
    },
    "HFT/USD": {
      "base": "HFT",
      "quote": "USD"
    },
    "HFT/USDT": {
      "base": "HFT",
      "quote": "USDT"
    },
    "HIFI/USDT": {
      "base": "HIFI",
      "quote": "USDT"
    },
    "HIGH/USDT": {
      "base": "HIGH",
      "quote": "USDT"
    },
    "HNT/USDT": {
      "base": "HNT",
      "quote": "USDT"
    },
    "HOOK/USDT": {
      "base": "HOOK",
      "quote": "USDT"
    },
    "HOT/USDT": {
      "base": "HOT",
      "quote": "USDT"
    },
    "ICP/BTC": {
      "base": "ICP",
      "quote": "BTC"
    },
    "ICP/EUR": {
      "base": "ICP",
      "quote": "EUR"
    },
    "ICP/USD": {
      "base": "ICP",
      "quote": "USD"
    },
    "ICP/USDT": {
      "base": "ICP",
      "quote": "USDT"
    },
    "ICX/ETH": {
      "base": "ICX",
      "quote": "ETH"
    },
    "ICX/USDT": {
      "base": "ICX",
      "quote": "USDT"
    },
    "ID/USDT": {
      "base": "ID",
      "quote": "USDT"
    },
    "IDEX/USD": {
      "base": "IDEX",
      "quote": "USD"
    },
    "ILV/USDT": {
      "base": "ILV",
      "quote": "USDT"
    },
    "IMX/USD": {
      "base": "IMX",
      "quote": "USD"
    },
    "IMX/USDT": {
      "base": "IMX",
      "quote": "USDT"
    },
    "INJ/BTC": {
      "base": "INJ",
      "quote": "BTC"
    },
    "INJ/USD": {
      "base": "INJ",
      "quote": "USD"
    },
    "INJ/USDT": {
      "base": "INJ",
      "quote": "USDT"
    },
    "IOST/BTC": {
      "base": "IOST",
      "quote": "BTC"
    },
    "IOST/USDT": {
      "base": "IOST",
      "quote": "USDT"
    },
    "IOTA/BTC": {
      "base": "IOTA",
      "quote": "BTC"
    },
    "IOTA/USDT": {
      "base": "IOTA",
      "quote": "USDT"
    },
    "IOTX/ETH": {
      "base": "IOTX",
      "quote": "ETH"
    },
    "IOTX/USDT": {
      "base": "IOTX",
      "quote": "USDT"
    },
    "IZI/USDT": {
      "base": "IZI",
      "quote": "USDT"
    },
    "JAM/USDT": {
      "base": "JAM",
      "quote": "USDT"
    },
    "JASMY/USD": {
      "base": "JASMY",
      "quote": "USD"
    },
    "JASMY/USDT": {
      "base": "JASMY",
      "quote": "USDT"
    },
    "JOE/USDT": {
      "base": "JOE",
      "quote": "USDT"
    },
    "JST/USDT": {
      "base": "JST",
      "quote": "USDT"
    },
    "JUV/USDT": {
      "base": "JUV",
      "quote": "USDT"
    },
    "KARATE/USDT": {
      "base": "KARATE",
      "quote": "USDT"
    },
    "KAS/USDT": {
      "base": "KAS",
      "quote": "USDT"
    },
    "KAVA/USD": {
      "base": "KAVA",
      "quote": "USD"
    },
    "KAVA/USDT": {
      "base": "KAVA",
      "quote": "USDT"
    },
    "KDA/BTC": {
      "base": "KDA",
      "quote": "BTC"
    },
    "KDA/USDT": {
      "base": "KDA",
      "quote": "USDT"
    },
    "KLAY/USDT": {
      "base": "KLAY",
      "quote": "USDT"
    },
    "KNC/BTC": {
      "base": "KNC",
      "quote": "BTC"
    },
    "KNC/ETH": {
      "base": "KNC",
      "quote": "ETH"
    },
    "KNC/USD": {
      "base": "KNC",
      "quote": "USD"
    },
    "KNC/USDT": {
      "base": "KNC",
      "quote": "USDT"
    },
    "KOK/USDT": {
      "base": "KOK",
      "quote": "USDT"
    },
    "KRL/USDT": {
      "base": "KRL",
      "quote": "USDT"
    },
    "KSM/BTC": {
      "base": "KSM",
      "quote": "BTC"
    },
    "KSM/USD": {
      "base": "KSM",
      "quote": "USD"
    },
    "KSM/USDT": {
      "base": "KSM",
      "quote": "USDT"
    },
    "LADYS/USDT": {
      "base": "LADYS",
      "quote": "USDT"
    },
    "LDO/USD": {
      "base": "LDO",
      "quote": "USD"
    },
    "LDO/USDT": {
      "base": "LDO",
      "quote": "USDT"
    },
    "LEVER/USDT": {
      "base": "LEVER",
      "quote": "USDT"
    },
    "LINA/USDT": {
      "base": "LINA",
      "quote": "USDT"
    },
    "LINK/BTC": {
      "base": "LINK",
      "quote": "BTC"
    },
    "LINK/ETH": {
      "base": "LINK",
      "quote": "ETH"
    },
    "LINK/EUR": {
      "base": "LINK",
      "quote": "EUR"
    },
    "LINK/GBP": {
      "base": "LINK",
      "quote": "GBP"
    },
    "LINK/USD": {
      "base": "LINK",
      "quote": "USD"
    },
    "LINK/USDT": {
      "base": "LINK",
      "quote": "USDT"
    },
    "LIT/USDT": {
      "base": "LIT",
      "quote": "USDT"
    },
    "LMWR/USDT": {
      "base": "LMWR",
      "quote": "USDT"
    },
    "LOKA/USDT": {
      "base": "LOKA",
      "quote": "USDT"
    },
    "LOOKS/USDT": {
      "base": "LOOKS",
      "quote": "USDT"
    },
    "LPT/USD": {
      "base": "LPT",
      "quote": "USD"
    },
    "LPT/USDT": {
      "base": "LPT",
      "quote": "USDT"
    },
    "LQTY/USDT": {
      "base": "LQTY",
      "quote": "USDT"
    },
    "LRC/BTC": {
      "base": "LRC",
      "quote": "BTC"
    },
    "LRC/ETH": {
      "base": "LRC",
      "quote": "ETH"
    },
    "LRC/USD": {
      "base": "LRC",
      "quote": "USD"
    },
    "LRC/USDT": {
      "base": "LRC",
      "quote": "USDT"
    },
    "LSK/BTC": {
      "base": "LSK",
      "quote": "BTC"
    },
    "LSK/ETH": {
      "base": "LSK",
      "quote": "ETH"
    },
    "LSK/USDT": {
      "base": "LSK",
      "quote": "USDT"
    },
    "LTC/BTC": {
      "base": "LTC",
      "quote": "BTC"
    },
    "LTC/ETH": {
      "base": "LTC",
      "quote": "ETH"
    },
    "LTC/EUR": {
      "base": "LTC",
      "quote": "EUR"
    },
    "LTC/GBP": {
      "base": "LTC",
      "quote": "GBP"
    },
    "LTC/USD": {
      "base": "LTC",
      "quote": "USD"
    },
    "LTC/USDC": {
      "base": "LTC",
      "quote": "USDC"
    },
    "LTC/USDT": {
      "base": "LTC",
      "quote": "USDT"
    },
    "LTO/USDT": {
      "base": "LTO",
      "quote": "USDT"
    },
    "LUNA/USDT": {
      "base": "LUNA",
      "quote": "USDT"
    },
    "LUNC/USDT": {
      "base": "LUNC",
      "quote": "USDT"
    },
    "MAGIC/USDT": {
      "base": "MAGIC",
      "quote": "USDT"
    },
    "MANA/BTC": {
      "base": "MANA",
      "quote": "BTC"
    },
    "MANA/ETH": {
      "base": "MANA",
      "quote": "ETH"
    },
    "MANA/USD": {
      "base": "MANA",
      "quote": "USD"
    },
    "MANA/USDT": {
      "base": "MANA",
      "quote": "USDT"
    },
    "MASK/USD": {
      "base": "MASK",
      "quote": "USD"
    },
    "MASK/USDT": {
      "base": "MASK",
      "quote": "USDT"
    },
    "MATIC/BTC": {
      "base": "MATIC",
      "quote": "BTC"
    },
    "MATIC/EUR": {
      "base": "MATIC",
      "quote": "EUR"
    },
    "MATIC/GBP": {
      "base": "MATIC",
      "quote": "GBP"
    },
    "MATIC/USD": {
      "base": "MATIC",
      "quote": "USD"
    },
    "MATIC/USDC": {
      "base": "MATIC",
      "quote": "USDC"
    },
    "MATIC/USDT": {
      "base": "MATIC",
      "quote": "USDT"
    },
    "MAV/USDT": {
      "base": "MAV",
      "quote": "USDT"
    },
    "MBL/USDT": {
      "base": "MBL",
      "quote": "USDT"
    },
    "MDT/USDT": {
      "base": "MDT",
      "quote": "USDT"
    },
    "MEME/USDT": {
      "base": "MEME",
      "quote": "USDT"
    },
    "METIS/USDT": {
      "base": "METIS",
      "quote": "USDT"
    },
    "MINA/BTC": {
      "base": "MINA",
      "quote": "BTC"
    },
    "MINA/USD": {
      "base": "MINA",
      "quote": "USD"
    },
    "MINA/USDT": {
      "base": "MINA",
      "quote": "USDT"
    },
    "MKR/BTC": {
      "base": "MKR",
      "quote": "BTC"
    },
    "MKR/USD": {
      "base": "MKR",
      "quote": "USD"
    },
    "MKR/USDT": {
      "base": "MKR",
      "quote": "USDT"
    },
    "MLN/USD": {
      "base": "MLN",
      "quote": "USD"
    },
    "MOVR/USDT": {
      "base": "MOVR",
      "quote": "USDT"
    },
    "MPLX/USDT": {
      "base": "MPLX",
      "quote": "USDT"
    },
    "MTL/USDT": {
      "base": "MTL",
      "quote": "USDT"
    },
    "MV/USDT": {
      "base": "MV",
      "quote": "USDT"
    },
    "MXC/USDT": {
      "base": "MXC",
      "quote": "USDT"
    },
    "NEAR/USD": {
      "base": "NEAR",
      "quote": "USD"
    },
    "NEAR/USDT": {
      "base": "NEAR",
      "quote": "USDT"
    },
    "NEO/BTC": {
      "base": "NEO",
      "quote": "BTC"
    },
    "NEO/USDT": {
      "base": "NEO",
      "quote": "USDT"
    },
    "NEXO/USDT": {
      "base": "NEXO",
      "quote": "USDT"
    },
    "NFT/USDT": {
      "base": "NFT",
      "quote": "USDT"
    },
    "NKN/USDT": {
      "base": "NKN",
      "quote": "USDT"
    },
    "NMR/USD": {
      "base": "NMR",
      "quote": "USD"
    },
    "NMR/USDT": {
      "base": "NMR",
      "quote": "USDT"
    },
    "NTRN/USDT": {
      "base": "NTRN",
      "quote": "USDT"
    },
    "NYM/USDT": {
      "base": "NYM",
      "quote": "USDT"
    },
    "OAS/USDT": {
      "base": "OAS",
      "quote": "USDT"
    },
    "OCEAN/BTC": {
      "base": "OCEAN",
      "quote": "BTC"
    },
    "OCEAN/USD": {
      "base": "OCEAN",
      "quote": "USD"
    },
    "OCEAN/USDT": {
      "base": "OCEAN",
      "quote": "USDT"
    },
    "OGN/USD": {
      "base": "OGN",
      "quote": "USD"
    },
    "OGN/USDT": {
      "base": "OGN",
      "quote": "USDT"
    },
    "OM/USDT": {
      "base": "OM",
      "quote": "USDT"
    },
    "OMG/BTC": {
      "base": "OMG",
      "quote": "BTC"
    },
    "OMG/ETH": {
      "base": "OMG",
      "quote": "ETH"
    },
    "OMG/USDT": {
      "base": "OMG",
      "quote": "USDT"
    },
    "ONE/BTC": {
      "base": "ONE",
      "quote": "BTC"
    },
    "ONE/USDT": {
      "base": "ONE",
      "quote": "USDT"
    },
    "ONT/USDT": {
      "base": "ONT",
      "quote": "USDT"
    },
    "OP/USD": {
      "base": "OP",
      "quote": "USD"
    },
    "OP/USDC": {
      "base": "OP",
      "quote": "USDC"
    },
    "OP/USDT": {
      "base": "OP",
      "quote": "USDT"
    },
    "ORBS/USDT": {
      "base": "ORBS",
      "quote": "USDT"
    },
    "ORCA/USD": {
      "base": "ORCA",
      "quote": "USD"
    },
    "ORDI/USDT": {
      "base": "ORDI",
      "quote": "USDT"
    },
    "ORN/USDT": {
      "base": "ORN",
      "quote": "USDT"
    },
    "OSMO/USDT": {
      "base": "OSMO",
      "quote": "USDT"
    },
    "OXT/BTC": {
      "base": "OXT",
      "quote": "BTC"
    },
    "OXT/USD": {
      "base": "OXT",
      "quote": "USD"
    },
    "OXT/USDT": {
      "base": "OXT",
      "quote": "USDT"
    },
    "PAXG/BTC": {
      "base": "PAXG",
      "quote": "BTC"
    },
    "PAXG/USDT": {
      "base": "PAXG",
      "quote": "USDT"
    },
    "PENDLE/USDT": {
      "base": "PENDLE",
      "quote": "USDT"
    },
    "PEOPLE/USDT": {
      "base": "PEOPLE",
      "quote": "USDT"
    },
    "PEPE/USDT": {
      "base": "PEPE",
      "quote": "USDT"
    },
    "PEPE2/USDT": {
      "base": "PEPE2",
      "quote": "USDT"
    },
    "PERP/USD": {
      "base": "PERP",
      "quote": "USD"
    },
    "PERP/USDT": {
      "base": "PERP",
      "quote": "USDT"
    },
    "PHA/USDT": {
      "base": "PHA",
      "quote": "USDT"
    },
    "PIP/USDT": {
      "base": "PIP",
      "quote": "USDT"
    },
    "PLA/USD": {
      "base": "PLA",
      "quote": "USD"
    },
    "PLA/USDT": {
      "base": "PLA",
      "quote": "USDT"
    },
    "POKT/USDT": {
      "base": "POKT",
      "quote": "USDT"
    },
    "POLS/USD": {
      "base": "POLS",
      "quote": "USD"
    },
    "POLS/USDT": {
      "base": "POLS",
      "quote": "USDT"
    },
    "POND/USD": {
      "base": "POND",
      "quote": "USD"
    },
    "POND/USDT": {
      "base": "POND",
      "quote": "USDT"
    },
    "POWR/USD": {
      "base": "POWR",
      "quote": "USD"
    },
    "POWR/USDT": {
      "base": "POWR",
      "quote": "USDT"
    },
    "PROM/USDT": {
      "base": "PROM",
      "quote": "USDT"
    },
    "PSG/USDT": {
      "base": "PSG",
      "quote": "USDT"
    },
    "PSTAKE/USDT": {
      "base": "PSTAKE",
      "quote": "USDT"
    },
    "PUMLX/USDT": {
      "base": "PUMLX",
      "quote": "USDT"
    },
    "PUNDIX/USDT": {
      "base": "PUNDIX",
      "quote": "USDT"
    },
    "PYR/USDT": {
      "base": "PYR",
      "quote": "USDT"
    },
    "PYUSD/USD": {
      "base": "PYUSD",
      "quote": "USD"
    },
    "PYUSD/USDT": {
      "base": "PYUSD",
      "quote": "USDT"
    },
    "QI/USDT": {
      "base": "QI",
      "quote": "USDT"
    },
    "QKC/BTC": {
      "base": "QKC",
      "quote": "BTC"
    },
    "QKC/ETH": {
      "base": "QKC",
      "quote": "ETH"
    },
    "QNT/USD": {
      "base": "QNT",
      "quote": "USD"
    },
    "QNT/USDT": {
      "base": "QNT",
      "quote": "USDT"
    },
    "QTUM/BTC": {
      "base": "QTUM",
      "quote": "BTC"
    },
    "QTUM/ETH": {
      "base": "QTUM",
      "quote": "ETH"
    },
    "QTUM/USDT": {
      "base": "QTUM",
      "quote": "USDT"
    },
    "QUICK/USDT": {
      "base": "QUICK",
      "quote": "USDT"
    },
    "RACA/USDT": {
      "base": "RACA",
      "quote": "USDT"
    },
    "RAD/USD": {
      "base": "RAD",
      "quote": "USD"
    },
    "RARE/USD": {
      "base": "RARE",
      "quote": "USD"
    },
    "RARI/USD": {
      "base": "RARI",
      "quote": "USD"
    },
    "RDNT/USDT": {
      "base": "RDNT",
      "quote": "USDT"
    },
    "REEF/USDT": {
      "base": "REEF",
      "quote": "USDT"
    },
    "REN/USDT": {
      "base": "REN",
      "quote": "USDT"
    },
    "REQ/USD": {
      "base": "REQ",
      "quote": "USD"
    },
    "REQ/USDT": {
      "base": "REQ",
      "quote": "USDT"
    },
    "RLC/USD": {
      "base": "RLC",
      "quote": "USD"
    },
    "RLC/USDT": {
      "base": "RLC",
      "quote": "USDT"
    },
    "RNDR/USD": {
      "base": "RNDR",
      "quote": "USD"
    },
    "RNDR/USDT": {
      "base": "RNDR",
      "quote": "USDT"
    },
    "ROSE/USDT": {
      "base": "ROSE",
      "quote": "USDT"
    },
    "RPK/USDT": {
      "base": "RPK",
      "quote": "USDT"
    },
    "RPL/USD": {
      "base": "RPL",
      "quote": "USD"
    },
    "RPL/USDT": {
      "base": "RPL",
      "quote": "USDT"
    },
    "RSR/USDT": {
      "base": "RSR",
      "quote": "USDT"
    },
    "RUNE/USDT": {
      "base": "RUNE",
      "quote": "USDT"
    },
    "RVN/USDT": {
      "base": "RVN",
      "quote": "USDT"
    },
    "SAND/BTC": {
      "base": "SAND",
      "quote": "BTC"
    },
    "SAND/USD": {
      "base": "SAND",
      "quote": "USD"
    },
    "SAND/USDT": {
      "base": "SAND",
      "quote": "USDT"
    },
    "SC/ETH": {
      "base": "SC",
      "quote": "ETH"
    },
    "SC/USDT": {
      "base": "SC",
      "quote": "USDT"
    },
    "SCRT/USDT": {
      "base": "SCRT",
      "quote": "USDT"
    },
    "SD/USDT": {
      "base": "SD",
      "quote": "USDT"
    },
    "SEI/USD": {
      "base": "SEI",
      "quote": "USD"
    },
    "SEI/USDT": {
      "base": "SEI",
      "quote": "USDT"
    },
    "SFP/USDT": {
      "base": "SFP",
      "quote": "USDT"
    },
    "SFUND/USDT": {
      "base": "SFUND",
      "quote": "USDT"
    },
    "SHIB/EUR": {
      "base": "SHIB",
      "quote": "EUR"
    },
    "SHIB/USD": {
      "base": "SHIB",
      "quote": "USD"
    },
    "SHIB/USDC": {
      "base": "SHIB",
      "quote": "USDC"
    },
    "SHIB/USDT": {
      "base": "SHIB",
      "quote": "USDT"
    },
    "SHILL/USDT": {
      "base": "SHILL",
      "quote": "USDT"
    },
    "SIDUS/USDT": {
      "base": "SIDUS",
      "quote": "USDT"
    },
    "SKL/USDT": {
      "base": "SKL",
      "quote": "USDT"
    },
    "SLP/USDT": {
      "base": "SLP",
      "quote": "USDT"
    },
    "SNX/BTC": {
      "base": "SNX",
      "quote": "BTC"
    },
    "SNX/ETH": {
      "base": "SNX",
      "quote": "ETH"
    },
    "SNX/USD": {
      "base": "SNX",
      "quote": "USD"
    },
    "SNX/USDT": {
      "base": "SNX",
      "quote": "USDT"
    },
    "SOL/BTC": {
      "base": "SOL",
      "quote": "BTC"
    },
    "SOL/EUR": {
      "base": "SOL",
      "quote": "EUR"
    },
    "SOL/GBP": {
      "base": "SOL",
      "quote": "GBP"
    },
    "SOL/USD": {
      "base": "SOL",
      "quote": "USD"
    },
    "SOL/USDC": {
      "base": "SOL",
      "quote": "USDC"
    },
    "SOL/USDT": {
      "base": "SOL",
      "quote": "USDT"
    },
    "SPELL/USD": {
      "base": "SPELL",
      "quote": "USD"
    },
    "SPELL/USDT": {
      "base": "SPELL",
      "quote": "USDT"
    },
    "SSV/USDT": {
      "base": "SSV",
      "quote": "USDT"
    },
    "STG/USDT": {
      "base": "STG",
      "quote": "USDT"
    },
    "STORJ/BTC": {
      "base": "STORJ",
      "quote": "BTC"
    },
    "STORJ/USD": {
      "base": "STORJ",
      "quote": "USD"
    },
    "STORJ/USDT": {
      "base": "STORJ",
      "quote": "USDT"
    },
    "STRAX/USDT": {
      "base": "STRAX",
      "quote": "USDT"
    },
    "STX/USD": {
      "base": "STX",
      "quote": "USD"
    },
    "STX/USDT": {
      "base": "STX",
      "quote": "USDT"
    },
    "SUI/USD": {
      "base": "SUI",
      "quote": "USD"
    },
    "SUI/USDT": {
      "base": "SUI",
      "quote": "USDT"
    },
    "SUIA/USDT": {
      "base": "SUIA",
      "quote": "USDT"
    },
    "SUN/USDT": {
      "base": "SUN",
      "quote": "USDT"
    },
    "SUPER/USD": {
      "base": "SUPER",
      "quote": "USD"
    },
    "SUPER/USDT": {
      "base": "SUPER",
      "quote": "USDT"
    },
    "SUSHI/USD": {
      "base": "SUSHI",
      "quote": "USD"
    },
    "SUSHI/USDT": {
      "base": "SUSHI",
      "quote": "USDT"
    },
    "SWEAT/USDT": {
      "base": "SWEAT",
      "quote": "USDT"
    },
    "SXP/USDT": {
      "base": "SXP",
      "quote": "USDT"
    },
    "SYN/USD": {
      "base": "SYN",
      "quote": "USD"
    },
    "SYN/USDT": {
      "base": "SYN",
      "quote": "USDT"
    },
    "SYS/USDT": {
      "base": "SYS",
      "quote": "USDT"
    },
    "T/USD": {
      "base": "T",
      "quote": "USD"
    },
    "T/USDT": {
      "base": "T",
      "quote": "USDT"
    },
    "TENET/USDT": {
      "base": "TENET",
      "quote": "USDT"
    },
    "TFUEL/USDT": {
      "base": "TFUEL",
      "quote": "USDT"
    },
    "THETA/USDT": {
      "base": "THETA",
      "quote": "USDT"
    },
    "TIA/USD": {
      "base": "TIA",
      "quote": "USD"
    },
    "TIA/USDT": {
      "base": "TIA",
      "quote": "USDT"
    },
    "TLM/USDT": {
      "base": "TLM",
      "quote": "USDT"
    },
    "TOKEN/USDT": {
      "base": "TOKEN",
      "quote": "USDT"
    },
    "TOMI/USDT": {
      "base": "TOMI",
      "quote": "USDT"
    },
    "TON/USDT": {
      "base": "TON",
      "quote": "USDT"
    },
    "TRAC/USDT": {
      "base": "TRAC",
      "quote": "USDT"
    },
    "TRB/USDT": {
      "base": "TRB",
      "quote": "USDT"
    },
    "TRU/USD": {
      "base": "TRU",
      "quote": "USD"
    },
    "TRU/USDT": {
      "base": "TRU",
      "quote": "USDT"
    },
    "TRVL/USDT": {
      "base": "TRVL",
      "quote": "USDT"
    },
    "TRX/BTC": {
      "base": "TRX",
      "quote": "BTC"
    },
    "TRX/ETH": {
      "base": "TRX",
      "quote": "ETH"
    },
    "TRX/USDC": {
      "base": "TRX",
      "quote": "USDC"
    },
    "TRX/USDT": {
      "base": "TRX",
      "quote": "USDT"
    },
    "TURBOS/USDT": {
      "base": "TURBOS",
      "quote": "USDT"
    },
    "TUSD/USDT": {
      "base": "TUSD",
      "quote": "USDT"
    },
    "TVK/USDT": {
      "base": "TVK",
      "quote": "USDT"
    },
    "TWT/USDT": {
      "base": "TWT",
      "quote": "USDT"
    },
    "UMA/USD": {
      "base": "UMA",
      "quote": "USD"
    },
    "UMA/USDT": {
      "base": "UMA",
      "quote": "USDT"
    },
    "UNFI/USD": {
      "base": "UNFI",
      "quote": "USD"
    },
    "UNFI/USDT": {
      "base": "UNFI",
      "quote": "USDT"
    },
    "UNI/BTC": {
      "base": "UNI",
      "quote": "BTC"
    },
    "UNI/ETH": {
      "base": "UNI",
      "quote": "ETH"
    },
    "UNI/USD": {
      "base": "UNI",
      "quote": "USD"
    },
    "UNI/USDT": {
      "base": "UNI",
      "quote": "USDT"
    },
    "USDC/EUR": {
      "base": "USDC",
      "quote": "EUR"
    },
    "USDC/USDT": {
      "base": "USDC",
      "quote": "USDT"
    },
    "USDD/USDT": {
      "base": "USDD",
      "quote": "USDT"
    },
    "USDP/USDT": {
      "base": "USDP",
      "quote": "USDT"
    },
    "USDT/EUR": {
      "base": "USDT",
      "quote": "EUR"
    },
    "USDT/GBP": {
      "base": "USDT",
      "quote": "GBP"
    },
    "USDT/USD": {
      "base": "USDT",
      "quote": "USD"
    },
    "USTC/USDT": {
      "base": "USTC",
      "quote": "USDT"
    },
    "VEGA/USDT": {
      "base": "VEGA",
      "quote": "USDT"
    },
    "VELO/USDT": {
      "base": "VELO",
      "quote": "USDT"
    },
    "VET/BTC": {
      "base": "VET",
      "quote": "BTC"
    },
    "VET/ETH": {
      "base": "VET",
      "quote": "ETH"
    },
    "VET/USDT": {
      "base": "VET",
      "quote": "USDT"
    },
    "VOXEL/USDT": {
      "base": "VOXEL",
      "quote": "USDT"
    },
    "VRA/USDT": {
      "base": "VRA",
      "quote": "USDT"
    },
    "WAVES/BTC": {
      "base": "WAVES",
      "quote": "BTC"
    },
    "WAVES/USDT": {
      "base": "WAVES",
      "quote": "USDT"
    },
    "WAXP/USDT": {
      "base": "WAXP",
      "quote": "USDT"
    },
    "WBTC/BTC": {
      "base": "WBTC",
      "quote": "BTC"
    },
    "WBTC/USD": {
      "base": "WBTC",
      "quote": "USD"
    },
    "WBTC/USDT": {
      "base": "WBTC",
      "quote": "USDT"
    },
    "WEMIX/USDT": {
      "base": "WEMIX",
      "quote": "USDT"
    },
    "WIN/USDT": {
      "base": "WIN",
      "quote": "USDT"
    },
    "WLD/USDT": {
      "base": "WLD",
      "quote": "USDT"
    },
    "WLKN/USDT": {
      "base": "WLKN",
      "quote": "USDT"
    },
    "WOO/USDT": {
      "base": "WOO",
      "quote": "USDT"
    },
    "WRX/USDT": {
      "base": "WRX",
      "quote": "USDT"
    },
    "XAVA/USDT": {
      "base": "XAVA",
      "quote": "USDT"
    },
    "XCAD/USDT": {
      "base": "XCAD",
      "quote": "USDT"
    },
    "XDC/USDT": {
      "base": "XDC",
      "quote": "USDT"
    },
    "XEC/USDT": {
      "base": "XEC",
      "quote": "USDT"
    },
    "XEM/USDT": {
      "base": "XEM",
      "quote": "USDT"
    },
    "XETA/USDT": {
      "base": "XETA",
      "quote": "USDT"
    },
    "XLM/BTC": {
      "base": "XLM",
      "quote": "BTC"
    },
    "XLM/ETH": {
      "base": "XLM",
      "quote": "ETH"
    },
    "XLM/EUR": {
      "base": "XLM",
      "quote": "EUR"
    },
    "XLM/USD": {
      "base": "XLM",
      "quote": "USD"
    },
    "XLM/USDT": {
      "base": "XLM",
      "quote": "USDT"
    },
    "XMR/BTC": {
      "base": "XMR",
      "quote": "BTC"
    },
    "XMR/USDT": {
      "base": "XMR",
      "quote": "USDT"
    },
    "XNO/USDT": {
      "base": "XNO",
      "quote": "USDT"
    },
    "XRP/BTC": {
      "base": "XRP",
      "quote": "BTC"
    },
    "XRP/ETH": {
      "base": "XRP",
      "quote": "ETH"
    },
    "XRP/EUR": {
      "base": "XRP",
      "quote": "EUR"
    },
    "XRP/USD": {
      "base": "XRP",
      "quote": "USD"
    },
    "XRP/USDC": {
      "base": "XRP",
      "quote": "USDC"
    },
    "XRP/USDT": {
      "base": "XRP",
      "quote": "USDT"
    },
    "XRP3L/USDT": {
      "base": "XRP3L",
      "quote": "USDT"
    },
    "XRP3S/USDT": {
      "base": "XRP3S",
      "quote": "USDT"
    },
    "XTZ/BTC": {
      "base": "XTZ",
      "quote": "BTC"
    },
    "XTZ/USD": {
      "base": "XTZ",
      "quote": "USD"
    },
    "XTZ/USDT": {
      "base": "XTZ",
      "quote": "USDT"
    },
    "XWG/USDT": {
      "base": "XWG",
      "quote": "USDT"
    },
    "XYM/USDT": {
      "base": "XYM",
      "quote": "USDT"
    },
    "XYO/USDT": {
      "base": "XYO",
      "quote": "USDT"
    },
    "YFI/BTC": {
      "base": "YFI",
      "quote": "BTC"
    },
    "YFI/USD": {
      "base": "YFI",
      "quote": "USD"
    },
    "YFI/USDT": {
      "base": "YFI",
      "quote": "USDT"
    },
    "YGG/USDT": {
      "base": "YGG",
      "quote": "USDT"
    },
    "ZEC/BTC": {
      "base": "ZEC",
      "quote": "BTC"
    },
    "ZEC/USDT": {
      "base": "ZEC",
      "quote": "USDT"
    },
    "ZEN/USDT": {
      "base": "ZEN",
      "quote": "USDT"
    },
    "ZIL/ETH": {
      "base": "ZIL",
      "quote": "ETH"
    },
    "ZIL/USDT": {
      "base": "ZIL",
      "quote": "USDT"
    },
    "ZRX/BTC": {
      "base": "ZRX",
      "quote": "BTC"
    },
    "ZRX/USD": {
      "base": "ZRX",
      "quote": "USD"
    },
    "ZRX/USDT": {
      "base": "ZRX",
      "quote": "USDT"
    }
  },
  "ms": 4
}
```

Find all possible crypto/fiat prices by multiple requests(up to 10 pairs in each)  

```shell
curl --request GET \
     --url 'https://api.fastforex.io/crypto/fetch-prices?pairs=BTC%2FUSD%2CETH%2FEUR&api_key=be05ce6a0a-f88ddd6bfc-XXXXXX' \
     --header 'accept: application/json'

```

Fetch one or more (up to 10) real-time prices for supported cryptocurrency pairs
https://fastforex.readme.io/reference/get_crypto-fetch-prices


```json

{
  "prices": {
    "BTC/USD": 70294.69,
    "ETH/EUR": 3272.98
  },
  "ms": 6
}
```

find the rest crypto/fiat prices one by one
find all fiat/crypto prices one by one


```shell
    curl --request GET \
         --url 'https://api.fastforex.io/fetch-one?from=KMF&to=CTSI&api_key=be05ce6a0a-f88ddd6bfc-XXXXXX' \
         --header 'accept: application/json'
```

```json
{
  "base": "KMF",
  "result": {
    "CTSI": 0.00809598
  },
  "updated": "2024-04-11 21:55:33",
  "ms": 5
}
```


