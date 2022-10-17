// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/bloc-chain/go-blocChain/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Kekchain network.
var MainnetETHBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://d3f63283928737711d7c381a69f9c8518f3a6210907716434ec635ba03333d101c7d9865a8384eaba914113f8678438a26a2f31bc47e41e5de57516cdaad77b4@127.0.0.1:30310",
	"enode://ec648d01bd85ef3d4252e2df90a2291ef287d1c4b6a7884d2e458b1a739cf4417cc4d5f1d9932b4cf86398d4c3598fd4f6066837580149137c2bec3c702ef846@107.77.215.225:30303",
	"enode://0646da1f72587cc4ff358533c50e8814d76a2eb0e94df58fa2b41cec9d6997417c39f83bb3106246bd742556d34f525c39d276e884b14adc69926f509c002594@127.0.0.1:30310",
	"enode://78e3baff0fded21bc6c798016bded6558435b7df6d7b67189a44ef62f346a889691389d1cdcd9849866cad2934ed4ef2b0d6266860f33efb1833bff24bd20058@161.97.121.84:30309",
	"enode://78e3baff0fded21bc6c798016bded6558435b7df6d7b67189a44ef62f346a889691389d1cdcd9849866cad2934ed4ef2b0d6266860f33efb1833bff24bd20058@127.0.0.1:30309",
	"enode://e30967a1533a6a3f02d9f7b9160f96a0f7f791084ef3c0723711fc4860c0e761401d8444ddf09ea11de3e68f8cb0da362269437d88128fb3af8ac9a59e0f1808@127.0.0.1:30309",
	"enode://e30967a1533a6a3f02d9f7b9160f96a0f7f791084ef3c0723711fc4860c0e761401d8444ddf09ea11de3e68f8cb0da362269437d88128fb3af8ac9a59e0f1808@107.77.215.225:30309",
	"enode://ec648d01bd85ef3d4252e2df90a2291ef287d1c4b6a7884d2e458b1a739cf4417cc4d5f1d9932b4cf86398d4c3598fd4f6066837580149137c2bec3c702ef846@127.0.0.1:30303",
	"enode://78e3baff0fded21bc6c798016bded6558435b7df6d7b67189a44ef62f346a889691389d1cdcd9849866cad2934ed4ef2b0d6266860f33efb1833bff24bd20058@161.97.121.84:30311",
	"enode://78e3baff0fded21bc6c798016bded6558435b7df6d7b67189a44ef62f346a889691389d1cdcd9849866cad2934ed4ef2b0d6266860f33efb1833bff24bd20058@127.0.0.1:30311",
	"enode://78e3baff0fded21bc6c798016bded6558435b7df6d7b67189a44ef62f346a889691389d1cdcd9849866cad2934ed4ef2b0d6266860f33efb1833bff24bd20058@161.97.121.84:30310",
	"enode://78e3baff0fded21bc6c798016bded6558435b7df6d7b67189a44ef62f346a889691389d1cdcd9849866cad2934ed4ef2b0d6266860f33efb1833bff24bd20058@127.0.0.1:30310",
	"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",   // bootnode-aws-ap-southeast-1-001
	"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",     // bootnode-aws-us-east-1-001
	"enode://ca6de62fce278f96aea6ec5a2daadb877e51651247cb96ee310a318def462913b653963c155a0ef6c7d50048bba6e6cea881130857413d9f50a621546b590758@34.255.23.113:30303",   // bootnode-aws-eu-west-1-001
	"enode://279944d8dcd428dffaa7436f25ca0ca43ae19e7bcf94a8fb7d1641651f92d121e972ac2e8f381414b80cc8e5555811c2ec6e1a99bb009b3f53c4c69923e11bd8@35.158.244.151:30303",  // bootnode-aws-eu-central-1-001
	"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",   // bootnode-azure-australiaeast-001
	"enode://103858bdb88756c71f15e9b5e09b56dc1be52f0a5021d46301dbbfb7e130029cc9d0d6f73f693bc29b665770fff7da4d34f3c6379fe12721b5d7a0bcb5ca1fc1@191.234.162.198:30303", // bootnode-azure-brazilsouth-001
	"enode://715171f50508aba88aecd1250af392a45a330af91d7b90701c436b618c86aaa1589c9184561907bebbb56439b8f8787bc01f49a7c77276c58c1b09822d75e8e8@52.231.165.108:30303",  // bootnode-azure-koreasouth-001
	"enode://5d6d7cd20d6da4bb83a1d28cadb5d409b64edf314c0335df658c1a54e32c7c4a7ab7823d57c39b6a757556e68ff1df17c748b698544a55cb488b52479a92b60f@104.42.217.25:30303",   // bootnode-azure-westus-001
}

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Kekchain network.
var MainnetBootnodes = []string{
	// Go Bootnodes
	"enode://4c77898a0f7dd6e1c3c0c1c33b29475d3b221cd35fa227213640e4b7f5e22d2e4f9ff8f686af7783916ffa4246a2a35cb8acac202672a7eded679a0bcdfc0fbb@167.86.80.61:64259",
	"enode://cc5cf9c7e0a4f60f90999fd3ae99d4af2d1d6deb001edbbf697139095aec48ffdb1022b47a63337b637f1e716c6c5f51c452ee4f6269ddaff3a9a1c7abd9eeb6@167.86.80.61:29384",
	"enode://bbb49a2c6d849a48ae522bff42d46fc18ac3d9643fd7b56374dd8778bfb724966faa6574d070e1ab770cb09b3e7b525bdbf0166b8fb38a4fcb53cf6271fd2001@38.242.236.247:64203",
	"enode://e6bd0a609407846fae37ce524eea03f3fb57e39876be53204eadc1903f69906f3b6e07c780ebb20b129a5ef33d10ee07a431f44619342e9080b645da2227d35b@38.242.236.247:64206",
	"enode://c1333dced1d42bb3e1a674eddb306f95c2d179402cb5bb183d9187fe31a08a2eef99c0e3bb1aec30812d26ea981038a345f29b22a4b1554973a8c8a6c31e1d33@38.242.236.248:12113",
	"enode://8ea70f12dfe1630862096735f3f422046fe82b6a966ef66c4274f918843f441007c9fed8b9d5b911c307bbf4479f020aaace42bdf1aaf174728e3582a3bb5b5f@38.242.236.248:12123",
}

// the main KekTest testnet network.
var TestnetBootnodes = []string{
	// Go Bootnodes
	"enode://e9788b0d60182bc65d22408599310396873c43afe8b2a8a27809b021851dbbbe1f0a8a952efb1558810bd5a61e073e0e86dd45621db7adbe5c02c4cb7c9adbb0@127.0.0.1:30303",
	"enode://e9daf5a7fe30e1a0128226a44500a3db1f8f9cb61c30ed4de6b9a0cc147b41059c1f262e53c5e8f234a439de5582f0302542ace2a25d22d511e78a0742fd12fc@45.94.209.226:30303",
}

// MainnetKEKBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Kekchain network.
var MainnetKEKBootnodes = []string{
	// Go Bootnodes
	"enode://4c77898a0f7dd6e1c3c0c1c33b29475d3b221cd35fa227213640e4b7f5e22d2e4f9ff8f686af7783916ffa4246a2a35cb8acac202672a7eded679a0bcdfc0fbb@167.86.80.61:64259",
	"enode://cc5cf9c7e0a4f60f90999fd3ae99d4af2d1d6deb001edbbf697139095aec48ffdb1022b47a63337b637f1e716c6c5f51c452ee4f6269ddaff3a9a1c7abd9eeb6@167.86.80.61:29384",
	"enode://bbb49a2c6d849a48ae522bff42d46fc18ac3d9643fd7b56374dd8778bfb724966faa6574d070e1ab770cb09b3e7b525bdbf0166b8fb38a4fcb53cf6271fd2001@38.242.236.247:64203",
	"enode://e6bd0a609407846fae37ce524eea03f3fb57e39876be53204eadc1903f69906f3b6e07c780ebb20b129a5ef33d10ee07a431f44619342e9080b645da2227d35b@38.242.236.247:64206",
	"enode://c1333dced1d42bb3e1a674eddb306f95c2d179402cb5bb183d9187fe31a08a2eef99c0e3bb1aec30812d26ea981038a345f29b22a4b1554973a8c8a6c31e1d33@38.242.236.248:12113",
	"enode://8ea70f12dfe1630862096735f3f422046fe82b6a966ef66c4274f918843f441007c9fed8b9d5b911c307bbf4479f020aaace42bdf1aaf174728e3582a3bb5b5f@38.242.236.248:12123",
}

// the main KekTest testnet network.
var TestnetKEKBootnodes = []string{
	// Go Bootnodes
	"enode://fcef32693af1c129c0608c2aaa260a644364d4b1289b8db41003a3a647b9c477df211d2622c218f0309fbe39e6ac373d8052009ad611d30c60accdc7fae4251c@161.97.121.89:19282",
	"enode://be46c3a4912307a6132e427732f1bcfd4459f75a45b370e369a4fa5843944de544468e7fa730b1af1115e37d6a54d14de4a6e8a4c2eb6eec86ab4a2c761e7000@72.189.4.201:19293",
	"enode://a48d80cd00d459fd2e3c31ad180d5109a1ec2a9e11d326591e5b4e163209b8cfea8192f8f0240aaa7cb598370c34ba3e07fd45d6acc537b68e57ee2e1ea05b00@72.189.4.201:30393",
	"enode://2501f28eafd9a065aa5ad13c955655e0c58508f0ab772eb19b51e9f68942631ba16e79a29c99e4ed1bc238ddd14d81ec13f50ac28151a7fdc5faa80061b2e807@72.189.4.201:10293",
	"enode://c3c5cf34f172ff16177321f2fa888a06d877b6fa73d952b7e5a905bbf8d32c9ae3b541d13ab82945a3e13aa281b4bbfb2e4fbb4ec2c505d8302fc4fab9f3a69c@72.189.4.201:30333",
	"enode://31877756edef4dab191052be2b4f9111d083734bf089ec70c6ce764a3c13d2a48a149d59c698e433de3d7e0b3fc786da0cfc4a8f81951ea387217a7c338bde41@161.97.121.89:19284",
	"enode://e2a8069889f5f867da599f52b80accdd369c1e88bdf322b4ee789744a9aa6dd5939b2372163a674bf52d4aeaa9f0c0aaa3ad95c7d8e29c5a61568cbfa855683b@127.0.0.1:64204",
	"enode://6fccc84880b89de3956a410a1a76ba57cdef46bf19f8a842017013eb3590ad2ed5eff29a8f403f76ee345430dc9bc71c5a00067e3e545c54259e8d38c219248e@167.86.80.62:64206",
	"enode://93dff98c642ccbf26f785b67695d86bf247381c124c120fea79576a35807d338e38df2ae0e9c40a4d97efd409e337dc6cc2ed94387d65685ffae1855e85c4929@127.0.0.1:64206",
	"enode://31b06c7f85b6891b21f56e728a1a1e298429e55b8896d3d67b088d61a738dd30db489c0d4f19569ac786c87916a1a2eb462a0d1527fe14ae212adb2788786918@127.0.0.1:64204",
	"enode://f942f6d4706bdcecb445d7d36ecc31ba3c377219e579b8f6a49675a0d5b58782cf16b82d756c5b0904280f11a459387f7b956747cd94a9ea96c8647a5f7b77ad@161.97.121.84:64206",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	// geth
	"enode://9246d00bc8fd1742e5ad2428b80fc4dc45d786283e05ef6edbd9002cbc335d40998444732fbe921cb88e1d2c73d1b1de53bae6a2237996e9bfe14f871baf7066@18.168.182.86:30303",
	// besu
	"enode://ec66ddcf1a974950bd4c782789a7e04f8aa7110a72569b6e65fcd51e937e74eed303b1ea734e4d19cfaec9fbff9b6ee65bf31dcb50ba79acce9dd63a6aca61c7@52.14.151.177:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://a869b02cec167211fb4815a82941db2e7ed2936fd90e78619c53eb17753fcf0207463e3419c264e2a1dd8786de0df7e68cf99571ab8aeb7c4e51367ef186b1dd@51.15.116.226:30303",
	"enode://807b37ee4816ecf407e9112224494b74dd5933625f655962d892f2f0f02d7fbbb3e2a94cf87a96609526f30c998fd71e93e2f53015c558ffc8b03eceaf30ee33@51.15.119.157:30303",
	"enode://a59e33ccd2b3e52d578f1fbd70c6f9babda2650f0760d6ff3b37742fdcdfdb3defba5d56d315b40c46b70198c7621e63ffa3f987389c7118634b0fefbbdfa7fd@51.15.119.157:40303",
}

var V5ETHBootnodes = []string{
	// Interchained
	"enr:-KO4QG-Ar_vKbp9mcfKngVYNhUemvsGWZqPZnqV6TSJLJSvxAxiU8f9kCCEQ4LZwsq_IoVkfpmazEmIXZS8WLvl4PIaGAX8JiF4ug2V0aMfGhMHs3gSAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQLjCWehUzpqPwLZ97kWD5ag9_eRCE7zwHI3EfxIYMDnYYRzbmFwwIN0Y3CCdmWDdWRwgnZl",
	"enr:-KO4QPgVZsM1GYKQDHgHmezVNAkWMpi0sGDJAUvFrrLfcfdEbIJPaDkCeZ0IMFfb-7N4knEfnprV1-jKCaqyQn3hA0mGAX8Jg7qpg2V0aMfGhMHs3gSAgmlkgnY0gmlwhGtN1-GJc2VjcDI1NmsxoQLsZI0BvYXvPUJS4t-Qoike8ofRxLaniE0uRYsac5z0QYRzbmFwwIN0Y3CCdl-DdWRwgnZf",
	"enr:-KO4QGtr5G_16en9l0qoJfdXK_wFvYsSol06c9C-Fp_CpuVzGd7irzBiL1ObdM61MqG8LAs2mBjkXnFOi4OoHuyjSv6GAX8JdJ4Vg2V0aMfGhMHs3gSAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQJpvrHzO4JfKqpzNZ1S369qJMsclda3b9sEIgTV865eEIRzbmFwwIN0Y3CCdmaDdWRwgnZm",
	"enr:-KW4QNeiA7orge3dofvFNKwrdm0eM3w7wlEqg10as8XBfa2AMcHGgQ0GgYCtgBUNlyoj9E9GjbZ7-zfZ6-cwp91U_ROGAX8J6c7Tg2V0aMnIhAJDl9eCE_uCaWSCdjSCaXCEfwAAAYlzZWNwMjU2azGhAtP2MoOShzdxHXw4Gmn5yFGPOmIQkHcWQ07GNboDMz0QhHNuYXDAg3RjcIJ2ZYN1ZHCCdmU",
	"enr:-KO4QOulVFMLfvgQAtyuBXGUbbWLdaH1eqW1Dbs2h9k3bKfGMqeWeXJ1iQi_v8T5XvqcO78FzkyUjZaMEUQAoSNHL_qGAX8JgqO8g2V0aMfGhMHs3gSAgmlkgnY0gmlwhKFheVSJc2VjcDI1NmsxoQJ447r_D97SG8bHmAFr3tZVhDW33217ZxiaRO9i80aoiYRzbmFwwIN0Y3CCdmWDdWRwgnZl",
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

var V5Bootnodes = []string{
	// BlocNet
	"enr:-KO4QIwbI1aN46Z_P931jojDJM-PsjRptlTxjUkTaH1O8znqbDtPtrtGZ1GSAwi4NlvSkvFS_2UZjHG-39ZVjGEwQPWGAYPA2PF4g2V0aMfGhNwalwaAgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQLpeIsNYBgrxl0iQIWZMQOWhzxDr-iyqKJ4CbAhhR27voRzbmFwwIN0Y3CCdl-DdWRwgnZf",
	"enr:-KO4QAYdw1SQi8o74v2KTnfB2vdKCcF1FmQWvrY3n-FEjRkvMKAQAesBhchxIXPov4VKJMlkceQeqBQvJIZPwyhSi2iGAYPA3kM-g2V0aMfGhNwalwaAgmlkgnY0gmlwhC1e0eKJc2VjcDI1NmsxoQLp2vWn_jDhoBKCJqRFAKPbH4-cthww7U3muaDMFHtBBYRzbmFwwIN0Y3CCdl-DdWRwgnZf",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "blocnet"
	case TestnetGenesisHash:
		net = "bloctest"
	case MainnetKEKGenesisHash:
		net = "kekistan"
	case TestnetKEKGenesisHash:
		net = "kektest"
	case MainnetETHGenesisHash:
		net = "ethereum"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".blocChain.app"
}
