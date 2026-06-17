-- sns_chain_registry import (keep .seedao suffix)
BEGIN;

INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x78f625a65fc316d32d98d249b698fb509a6d98f2', 'tanghan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8c913aec7443fe2018639133398955e0e17fb0c1', 'baiyu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x82944b68bb92fa11764041aa61204b5fdc85f429', 'frozen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x27b503fd677dd889ea3e7804df492d2d809e6052', 'shawn.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xab206fb4b72009c5c7e78c2a91b0b8e63fb73e34', 'cybercaf.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x13e283646f203ef6b24cd5c8ed5dbbf29dd0d39f', 'rebeccawong.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x52ff0010d4058bb726725ad76f542c53033d67a1', 'pm.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4ed2f7eadd13cdc339b67f371610bf26224e4b98', 'chefkang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc341683d1ee57b13ab90a1fc00f38dc6f7afdd14', 'alanwei.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9d6b1a15d476bfbbdf4274ef6b405086cd7258f9', 'joanna.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4564d5a8bb409272f1fb4ae4c8b45fc0eafd709d', 'hzmangel.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd85c413da833cebd8338138ccefa04979df70e8e', 'wendychaung.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4d4b78d37090ed3e1eae6779ba2c3d6728052915', 'amanda.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9672c0e1639f159334ca1288d4a24deb02117291', 'kc.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4d6654beec5942aee8e3102528a58fb34a0f1b19', 'grace.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb8aee482305b7b219704350786fce47847cec996', 'lizi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb8abe31a6caf1217d632190559b2c1fd909178ec', 'suannai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x290aa5081d6298c181949946893876a1b76bf144', 'jialin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9fbde1512d2aeb2000866fa6f8a01e6bc80559e0', 'jack.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc92a88e2d17887646a40ddb0ddbbbe5c2a72512c', 'fivea.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4ffb83a8344f33453fbebe5e126534407ceb32f5', 'jason.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4e5e4a07dd21e4b840052d98e113aeae5798da3d', 'yao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa50bb794012b9617cdcac8653f8041842f3c4315', 'xi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xefd67615d66e3819539021d40e155e1a6107f283', 'hyll.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x332345477db00239f88ca2eb015b159750cf3c44', 'forever.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x42b81011a367d3c7f4e4b570fadd51209c93f287', 'alex.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7d53ba1d70a66d572a90e3e80ad4253769340ee7', 'munan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdafe3083bd4865ade8fea183c0e058a585bab68b', 'hal.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x043aee850680c4fb982761b585929834c26cc32c', 'kyle.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x796ac058028f19974a4f8256f6b30fedcc96537d', 'ache.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4a40eb870dcf533d4dc097c3d87aafe9f64490a1', 'icepy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x85633c25d2518dc5e6e05b8c9dc4596db2918575', 'waite.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x608a9358354a93581e686ac7ffde9d5fadbcb7e0', 'vincent.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x29faeb62f4e9e1456c2feed5c6f27b13fae91dae', 'owen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9731c556a8f285fa33ff3dec4add14784b944036', 'walter.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4fa96d31dd15ce27537130d4b56caee0cb436051', 'kim.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x97d455835a3ae0f0616f280286439e248f8cb030', 'zoe51.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcb0ce7fd46bce2ab4d561642e601ff14a9220422', 'martinlok.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfa860889ee6e3c715a5b76d998223df040b167da', 'bubble.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x625b261e91efd5ec12ace7a51d56b52b5040e8d9', 'somehacker.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x35fe0aaafe30b98c90d65a24590ae520b516cfae', '0xhaimian.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2e7e94be872319a0d991f704e1f4fad385f5855d', 'victor.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3caee6446cbbb73569ac230939b4310dea505fa9', 'darcy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6bd68d667eccea663b2a2bee1311f7ad6c817602', 'walton.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9bd8011af154f63589d8c71d50238c44ad7a6764', 'happy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x009c35a834a1e94557c772c8719925ed245520b6', 'leon.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfdaae168c3d835da1f974447cdccf8c7476d0792', 'black.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfc72a128cade26a6d1d06888b732ebc74762e6a4', 'double.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x13bd0c79ed0860dc38bfee40d07f49449944a9ea', 'sed.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xac85db17b614c94332b494ee8329c308f395a053', 'davion.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf62aadadf80f0cee151af3851a9b3866255c6cd4', 'nicodemusdu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc077cdf746717f3e576e91809e72fd106aff5e52', 'smithereens.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8817e259b93256bbb86de6cd2e9b8612763cc674', 'usopp.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcea5e66bec5193e5ec0b049a3fe5d7dd896fd480', 'caos.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9c5ab52309c486ff93b20772ec4b113da3b42017', 'davefff.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2980ced2b9123d36b26c6ae938594ab3c5bb465b', 'deyson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x88a0093c293113177175d3a995c49eeb68a887d9', 'chj.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x597f0fdbb96485cd258777248052ff0e4781fe4c', 'tovarishch.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x10b4bd1a211bdd057ea16a2e13c086f67b587b4b', 'manman.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x47a7202a726324788609003549675f1dbc645b07', 't1ng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x902b13e46305e3ae10f70871c6d0bca9e5446b07', 'guyi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9a9737e821bb216ef74cd0e6b2dd3e4fff98c681', 'cl.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa912cf5d10c87b7fc9cf6b47d203bd1052d02c67', 'sarse.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x79eed13cf2d1530e36d07b9a96c17ab795302a54', 'yelei.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe7339d4421c2d0290d38eaa06657e9e2dd8d3bd8', 'sunzhe.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd847ed72649a39aad263ac7a1dcae6eeb9d51e16', 'yzbban.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa76a36765c8af45affec2b5de2abfc33adc1488a', 'yuy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3f13c31474141f6912eed296220c0996c90253d2', 'ines.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x41d2ce62cd81d9ccd5c6890dcb44267b26165f85', 'xs.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x103af2b33400a4f955744a7f70e749435444f808', 'grantyu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x10ea167893cb51500ec829222e3d17a968feb650', 'henry.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7a6ff8906f68c0ceafdd61c2e5d39834cf0b39f1', 'queeny.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x41448e812733c10b6e74a90ac8086ba3a0766188', 'beilin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcea9f52fb205846b0c2956ab347e078b564277f0', 'wonder.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6cd5f918c5125f4e421c354e8b27d3bc9ff567f5', 'yohji.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0c05aacc260d5ebfccdd9b06ab3633120cbe9676', 'backdoor.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1701bbf4c5cfbeeebd846866f949001809835ff5', 'yang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x65d5b68a7878a987e7a19826a7f9aa6f5f92e10f', 'demian.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5ecc0ea06cc6ccdcf52aa0d1139c2813bb6ebef4', 'pb_lookme.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xabef0820004c4108d2cfca0b8cf326aa3f08eaff', 'chunyu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5e66f491a1caed9291b59a6297f1a9dd71294e05', 'suetang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x20669775d663d9a461ef2b2efca8fe43954f4b13', 'jerrypp.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd8652a853f4ac24c797faa02e4633bd69ac9394d', 'philhuang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x42b0765b70af2b6910f6281b3dc0c34a9cdbb65c', 'jomosis.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7ff4a502764b47a61063f551603535d8b8c91a20', 'helloworld.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xee139e7836ccbeb672a7725aca970cd5af8bdaf7', 'kitty.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x909085925b6321e64f00b5b89a7604fd1dc2754b', 'ivyfanshao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcc8a60d940431f6a1e012291da19e8fc7d183e9f', 'whitepaper.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe3a69e461c2502a6685d8e108366ad00f3208225', 'cryptoc.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x27aa5a880b512beaf40d1b2e0f256b09f5711444', 'robin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x991b1ade06d49454fd7a2caeaf31f908e00d1c63', 'nirvana.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc32428b4b31873f41e6a6b81028080469e2d4492', 'seed.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe30ce0ad68192cf1fa258df0b7d7a6fcf64b580b', 'rosa.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x373d8d85ad0da98adb4c14c3ec93b35eeae1f9c7', 'gaowei.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xaaf2f3a917e5ac907efd9c0f3fd752cc9d66dc08', 'ashley.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x64332bbd3735eccbbf7c62fb66fa32cbb0a9279e', 'cipher.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6d3fa2b0faf3bc37874b24c1b1df9735cb10ef00', 'chiangmai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5eb71caeadf22107a578bc5261ec9a1e16fad5f8', 'ling.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x01974549c9b9a30d47c548a16b120b1caa7b586c', 'blanker.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5fe724baa783183d565bfa32bbf1e3daa8e845f3', 'joyid.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3f7a3df9c6f5574b6e67fb26955b75e52a4e2b7e', 'alen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdef62635d2f06e6ee32f45e49bd1f3f26ab20f3a', 'johnson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x68054004024e88d1de42dc8aa5bcea9afaddde61', 'polymorphism.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6bfc3e4dabb4b7cd020ae3d7ec42590f03d427d4', 'xiaosongfu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x92d27113ec390579c24654ff340ca5d2be9b32df', 'suuuu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x18f958a14c0fb4f38683d23088082d21435f7f0c', 'apple.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xddf29c0e97601d8ab661c4b3c3a8ac7d9bd5cb77', 'rekko.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x34cb6cf059a80da45d2e76ddf6d5d1d3a0fd80db', 'phoenix.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x97dbe1bb349dfd563b1c7ffe70bbd73c0b779cdb', 'purple.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5b228fcb60a95f7e749e00c4b88154fc17e706dd', 'wenq.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4adf647e256b9b1e8086064437f6b404a5c7b421', 'xiyang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1c04e9363fc7751dc237d6eebd71533514f914d2', 'eric.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x89c60c01f2e1d1b233253596bf1c2386bdfeb898', 'arjun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xec5041fbdc10166717d4a84287b02a6bed152e4b', 'gwendol.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfc0473888cf4d8e7c291ab0e9e6c011a3bae516d', 'metabeing.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x132e925e1f13e3a2e6f948e8a0ad0a1b99454af8', 'xingtian.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x05bbaabf7dfbcf371c3741737c8f382907399a24', 'wukong.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x635bdec4531c95b0f4f761cdb5146f86e5ccc015', 'satoshi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x671215a06a0a4fe039a7ba8a282d178512929572', 'julie.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xeb697d3b2ee49a86861543f41f745e13b7785a3d', 'zoezhou.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x36764af31381057e722a003a69d089d54fe2c61b', 'chaolink.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x098f2c764f892c24b9abb9656761eb5946370b35', 'shawna.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xab48bbf712d64abe1d0515d0b6c0774414cef3cd', 'helper.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfb6c83719ae7ba42ecb30dacc40839fb0c96c418', 'mirai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8429837a4d311a5b82215fee1f96bdfcdea9cd73', 'suki.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xae5be9ddb562d34b9e54f5ebbafd2ef787967c52', 'jessica_feng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2fcffdb455087ba3490c72f7a57890f7fca34e05', 'seeu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb2206f941a5206b55bfa33e665687523a8bced80', 'swiftevo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x03e2c10b74a260f46ab5cf881938c5888a6142df', 'liubing.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1320b3acac78eaf26cbf01697efb65afef320da3', 'zhangchao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9bae3afe6605658b029d8da152452a8e32a3f116', 'shaoxia.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd767fb7f40209473f3cf1052350e8ed2290514f7', 'lyle.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb18e3c41faf4139b89b4ebf1f5ef645a3ad0ec7f', 'yyds.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3f07912e55e96fb0fc0cacb38e10be1d58045228', 'jedi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1496f0bc152842c79993946946626af285b2b077', 'momo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc1c2d0860a213c630c7e285eb5b8b1ebdfa205b1', 'ticko.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdd8293f05fabfda054f249be7d61f5ba56169d11', 'clark007.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1adb86940e7c50fc6b82765cefee500feac5937e', 'king.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1f612f3c6250e1de6e68d8a5ab393c91eb260b77', 'empty.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3b75b768dd0dc2f71ac0bca8713ed5e9c82cfcce', 'epimetheus.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3e2141d27c11da046e443c53441ed53b390cb15a', 'young.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5edbb5bfbfad5e5ba20e699dca2452aae2a8f55f', 'nick2024.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6c7c73aef66836a1eeb524376e13e71cfb52d3ec', 'axin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x68690c4cecf4245e09dc7e47148b6eff70cb7061', 'clare.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4c1f5e828ab813ab5f81de61231a870f931dd292', 'ordi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb177861050d607ada9ab10886b5273a021e9ee2e', 'xcsun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb20dc89ae3c6c960234169ea674119a4011a34f3', 'xing.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcc5db6f29944979a3f822e705412b471dcd9da74', '0xmoyu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb789a1297b1f5d46435d1679e18696438a795b04', 'soda.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc767e00541f50337424463f8d3ae1a04c478c00c', 'tong.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7abf99b7e65a43d2af04e8e0d47d3361494513d9', 'btc123.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd2a314a059a2fd576e906c1bbae559f8c2406b59', '0xbensha.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5b5f3ebe7d9e1f3cafc22cfd03d50805dd0b5b68', 'alex_lee.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1b9eae3d365b16f36c3c3b197d216c459dad3b72', 'nicky.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe02533f1b8bc4befd8561df63721a96936f978f9', 'lee233.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa69247869a9a7387ba07dee89110f189fd225993', 'heihei.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5d3a6d7a1419f7f9b61b52d3daa6124742ed035e', 'chase.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x25f89e72c8b362ba2063b72a60f27acf76e37406', 'tivita.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc163a998f64efaffea83cb026d1941e43b13a9ce', 'lodar.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x750ea21c1e98cced0d4557196b6f4a5974ccb6f5', 'qiaopengjun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7867450c3990237cbe34628f1d872e8cd8121220', 'sunrise_feng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb39bebc396c0849d1f894efb9037b76640fff9ce', 'amagi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb878d813495c6e42bbf8fdb6a289d3622d06ebd7', 'ricardo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xec17f69034005b5fa8ded070523952d8b7d4b37a', 'marlboro.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd4302bf70541bafe3ac46bbfc3c7eebe20ccf7db', 'vfan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4b162dba01943bd8cd8af668be72fb769999ac83', 'tony_wang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xed3fb32c5c61f00f4dc7f4b64fd8378bf1a15f61', 'huan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb6de28dffb4457ec309a76ec29e268798d163f2a', 'ekayana.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x057e21bf3f7592247bf8ffb7c1e833ac553bc6d7', 'gsimt.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd287b3f7f423d4fd7f3c000cf89f8f78fcd2a8b3', 'metahao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd12e8258a8dc23c75484e6b037c713269060ad63', 'erica.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdeae2409aeaf60fa2bdec5cd818deabe0fd17662', 'sethpanda.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x32a6206c3260ac03170b4d810b1348404e4d938f', 'chan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfde65ea7a7af9b4fecd585db65e0b0394a2a55a7', 'welten.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x77fb99be6caa5e2b5f41c4c680a622ec66521d28', 'clyde.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9c017079fcb2fd7f38ebd41ebedcfd86d2d29d17', 'kylee.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc052261da7602245558b297c587a8545e67d1109', 'bitcrab.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xda21996423514f2d83cc3c271c562be47a22d869', 'nature.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8bc5c516d460afc3ce001fe4e0a9f1faa424b81d', 'jessie.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9e89a144f7c05eb0d258b61a71d00bca79e42715', 'sirs.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe9c8162761e08434993f9280b285446633d23297', 'fuaijun1.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x800819b08c66a1ecbea09e4b64c7e3412c104df5', '9999.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf9a35867828417d6b42cc66c004126ea98083e28', 'westwing.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd1a5b91957530e1b3e9cfac1543467c60c352f69', 'root.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xec1e9d03a08d3d2373666d9606313f4901b4147c', 'token_wizard.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x163052a1efa47c7acf9f9b213df508dc3d3ebf2a', 'ode233.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcf1b6522890c640dcd5075284d3d73e193d21839', 'killpanda.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x07888e3b3014e364a6738150bc6072e88113df6d', 'laura.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe9a7a20ca704f23bb0979e6c6138f85fc1d93336', 'lalain.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x73b92634910d45412cdecf7a8d9baf83a17007b4', 'sumzero.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf490116431d6d930ab27039a2bd92121d8fcaea7', 'obluda.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe6fc58e4466ddc40d2d86fff441538be4320a8f5', 'breaking.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1c92851768b244a121384f55c5a5f409a109038e', 'xxxl.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa21a371abb112add73eee2fcba871c97642277ae', 'xxxx.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1a2b1ac310f34970bd110a00a1a7a0d411aac501', 'cuckoo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe8a9c8b1e40e59485978943ea12567a3127533c3', '8888.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7b2aa84f57fd4c8a9e5b2414821eedff53173d3e', 'kaori.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x443d40b2aa10cb01884a97281dc8ac3ccb0eef2a', 'fifteen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x50a6b6afec8f9968ee26aa55f9e34e1207750a03', 'love.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd5ce01d14f2ca6f8d061bc375334e8647a9474f8', 'koo_oo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x51ce2507f56bc46c8c31979b8724fb084de3a686', 'chiyuan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfb3bad2243a6462504a26e13132a47fffb81fa27', 'civildogs.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9bbd28844ee32784ab1e00be4d9ff144aaa826c7', 'pixel.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x38c92927806af571afded66576f648146f691d38', 'luckykent.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf9c5525792df4d70c287b872bea21c2e44df87ac', 'blk2.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8c17dde9d8bb4cee8fcbb7e135afdea01e8f6b83', 'bitcoin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x475186d1100505e68da4a4c936a8ce082978cd3d', 'leodiego.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb95cf73b8d2bcc0f9ae81a513d6aa465dea493eb', 'hank.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7447debdacae3638bdbac2f4a443c5615d3f007d', 'handuo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xaf53261cc893ebed1d2c688c64a2fd26b549a816', 'mailajinbao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x895ced1d75fa2dfd159c0242c934af645ee535ba', 'ulysses.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7b42da12edcfec3634cb518c2f8e153b88595530', 'fhxl.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb8750af88d2e95d752b6dded0592d7e12c907aa1', 'epikhong.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbecc3ea31ee6d12c81b01cdedfb75bc3dded9f34', 'laika.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x97f44939e57ba2341ff10dc25cb4314c783539c6', 'administrator.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc8d665f8453ebb6e0408f13bade8fbe90c87b2d3', 'dice.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1b76b011c0a1e03896bfc8923be228e430bb9049', 'nechi_luo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x60b7eca6856ca27d409a375fe0b9d7e212975cd9', 'okay.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x70c58cc68eec54f9aa2694aa5575b6ac80da440c', 'aranna.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5f5916e8c27c1dc6d99602767d4df87dc04abe7e', 'georgehu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x40adea030bbb77f450fb19bd6a5512d51a0ae47f', 'louislee.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfa483c6f280782714813c484ea05bab55e9f81ad', 'alan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4c55c41bd839b3552fb2abecacfdf4a5d2879cb9', 'pipilu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xecffb8de150a879d579c58f004727e92cea14ff9', 'melo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf17a1ae74a192af85996772aea3e31ca1254ab60', 'master.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa21ba8ca88e3b8942ed419ea4c44a864c9d56bb0', 'ruru.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x02d4879d7997ac00469a10cb4fa407dd2907c752', 'yuli.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbe3018ec0a8d780833336251e591a58097d8618a', 'dyson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd88dccccd6e1c614614e21183c97e5391718d929', 'rectinajh.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6a9725942d00f7c0d6b0ed1dd215990a653d3813', 'jan21.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xab24425162d4667db98f450f4208c13ee0278815', 'judy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x78c2d5f600fa2fb45f98b279df87b1234678985d', 'ellie.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7789101d739e292ec509c824f54216124f979824', 'ericnomad.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1162201a5510b21abb12ce49308b3db511aeec1d', 'flower.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1cff3b7b4ed8b58ee756af8aa54ccfdfa3928ba3', 'platon.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd245540c2eb89ab9359127b5d7d85e02bfd815a4', 'crypto.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2a174cf81204e9c7fbf7cd260871af4f5b0a482a', 'sats.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7d3ed7e4798b67e02b99e772c453b98f4125aff9', 'henrybewell.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc0a129c97292d9a49fa5b80e970862588331ba0a', 'nour.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6878b312a886cecac9db1ece5b6478e2ad1fcc13', '0x99.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4689f7f66b968b0bd9d6caef5c1316f30c68f4e2', 'w33d.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb8372df9d06345eb96e27d0ec87341b10f23a6c7', 'kaemi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbfd8cba6a1e10e1ab4fa11a0062f4e52e13d260f', 'huawen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0cd131d695c0dd47648e70ede3bb13bcadea6f40', 'niraya666lzy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8068bd398396b060d820745fde9c545a38bb776d', 'baby.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0bb9220b5192d417aa6036541eb0b8b435d08ebd', 'yao24.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1e2c17a74d2b9f602bfa7a572811883249450f21', 'miosleo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x63927a83e862588f1374cc8303fbf5b06720602b', 'geemukden.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdd41cf6ceecf0cd3779c6eef6ecb1cb68fa60fcd', 'saturnliu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc0ad247d579e11937486eac762f708861357101a', 'noah.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd59f74e988261683530f2e31e2994019906cd633', 'chauncey.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe363e25499464d1156caa2237fd1a5d3934cee7a', 'luckycat.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb39d8381a86a713d6944c49fad8ba7fba4fc61e0', 'khli.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2733653946a7bd78997160d17dff066026f96428', 'exeleo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5f83db8f64d698444be98d0fd9f264a38027edbb', '9952.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x43ce56f22720a07991a36960abc7d7649ab4c90d', 'jianru.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x54a0edccc4cd1d821eebcbd1cb79e90a5ac5db36', 'crypt2miner.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xeebbc090f07a34a9b69e655d287da66606c59150', 'china.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xda8a0fb4fd57e12c1e6623375a9631dcd6400c28', 'verlo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x98a26821527ef9aa298cff42a06bbcc6a6575977', 'geek.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3a89de26aaec8f6bcd841c059e6d674f8f88b382', 'huatuo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xaec2e770042682bd086e27a87a1ed2d75a5c8d15', 'moun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbb8b4d64f99f714954522c3bfd21b1cb1a6de9d8', 'mrly8621.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9c4e2b80e7107f8ba8e835f42119483b7763d74d', 'overload.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0e013ebccd71ccc8ae546cdd1bc254b70cd904b6', 'trancongquyen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbef951e50c30fc7d6e0f8d1497ec08ca3c6a7572', 'wangxj70.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd9897e8b7a587ea154691199289f2341e0a1c20c', 'zhangqun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x24c5b0815fbed59047ef2af8537947b758af2945', 'edward.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1fc7b6c13684d65124bdc6793bdc0a16e593203c', 'sanshi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfcc836c2bab4787a4556c3f01824c5338ec778cf', 'null.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfacc33eda2ea43d7c7e00343cf0ff119e43aab49', 'denser.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x76ae1e4e0d17dc32b0fe0478f3ed27c4fe45d174', 'feique.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe95061f332647a784570272890df2cc6ea76a91f', 'feiqie.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc8870ca4a8e3f581821384f3c927a9494a661b35', 'llyy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf8690b1ba8cae43190cac09eff23b0763bc670fc', 'hanyang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa7c4a6b1510dfc72b55ffcc9b8152cdada091489', 'yunfan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2b40590f525b4f552b0efb85333877b4a0769400', 'zhoumo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xce18a7452d57579d6c6b885b69b1978067cd2d87', 'mona.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb3ded5225fe4a42766fee4df4753b020e93dadcc', 'seven.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x934061a4b54bd11f5f6aca7f9373bfbb74658c0f', 'flood.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc81d59059ff38d61acd50bf4421cfa5337544790', 'nicorobin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x940fcb5bbcaeaa1b434ee85a1ba908ff1e3d672d', 'terence.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x615b95fefd5307964de2e38d544bb835c5b79717', 'emma.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x04425df955de991811735da40fa5cf8ed570ef58', 'vanvan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x70fa41ec7506aaeacebf04eada4ff9fd496344a3', 'will.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9dc8cba394ab678d4a8355753870eaf0d82cd773', 'layer.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x77cb545674f583ebd0453459d4b680cc4e19e0d7', 'johnny.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb2e233966f32bf5dddc3efb61e22963e9e84c88f', 'damien.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x332678d58c224292d440f47684812dd7b9fabfa6', 'tangut.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfce7d57ffed117ffaa1ec3d4c0b6154e8384cc06', 'ricky.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x75b90dffbd7c75c42d1ef9513ff9be66806fe232', 'web3jv.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9ef40f93886f31bec12eca10f960832565b9cf43', 'jumpan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbf3de8b57cafcb97d17341a29d932e77e56cf171', 'jincheng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x43492e8a862baa9c9fc505d95b9e67e21622ff7a', 'marcus.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd6f6a387019961592c97670fc4c243f5fc6d2153', 'tonyjai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x38f2a3b2a83c90e1ed30b07badd59d3c14fa97a4', 'lllu_23.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6287428f0b97b0b7966706014b8ff1f34a93fae9', 'kate.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8184561204f6180496bfee4ddfd9dc2c0b57b5de', 'keiko.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x610566f55e26c87915123120253a2b7a23daf750', 'runningchild.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe7fc2522d0d3f45657935fa6cb85771bbe02f6dd', 'chainyouyou.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8ad7143a1da425e961673be32744d0faf254b8cb', 'bibiruyi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x38f5df039b0f0fac72523eb0dea196c0e6811d4a', 'raven.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6098e17a2945d2858368329e03c0a4a64a549814', 'bobo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbc060432d484a50c9822c6992118424f718430e4', 'mr123.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfcd05fad2257b08491a023793c4e5922f4df6159', 'ekkoyang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x913d7434a18f0ef40181b92a9603e0425ab0addf', 'xiaokong.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf3bf6bdfea27464840d53b07c33ea21c951d3746', 'hotwolf.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4a449ff887e534ffb98016d38bbc9677bf4a755f', 'uniceros.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0db515dd376e4b0d5e747041c04455ecccddd5cd', 'kelsen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa7f8bd4a9d966e19890e62ea69a8f4ddce083499', 'ceshen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa8273c123fcf8728b95ab3b4b4222c426c6c1fa5', 'patcan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x74b04e4d28a8dcf21b481b384f3023b59dfb16e4', 'lauragao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x71856917bcf6ba1bd0f3e53ac9adf475d5a6de70', 'braavos.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x89c522b9f73f46985baf4cca047088aaf561b816', 'hisky.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc8f11e8a20a4c0d34509d89ce9e32b53c001076f', '_0xmilky.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6a09c478559565b33849eff37e26c455bb61a9d0', 'ethan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0c70b20eef88e15f601cc26ae5df4d038bb293ea', 'wendy112.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe2301dabba35e886f4ec121f13abcadf4632b90d', 'linlin_g.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1a709c17b2454211dd573889075bfa3b93dc847e', 'periclesding.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf8279bf46bf0a5ca756c84350332d987d6562950', 'txxx.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xae0e344672a848a174aec37ce4144e05b80c440a', 'jennygao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x787db9b5e3ef08d24c1ccc0943935e93fe9efe0a', 'sevenfloor.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x66042381fe037f32557eb68ecbaab9b5f394f5cc', 'ergatea.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd82f8ff7535330e300ba59af61d7acd1bd6b35a1', 'xiahua.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc95dfe43cd56385a96fdbf6c1f80e15a83e922fb', 'carson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0505e7fddf5a88a8a3743ac579cb6f33c181d4bb', 'tuntun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc19be75b8b9152d884987e1b58b3f18a94875396', 'liyang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x296e457480e174e72fdb4779e70f1c3335271027', 'senn.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbe8101847d9198a0c150cfd681986ac963ebb1b5', 'miuyun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x632807db968a4aa363514c6cfdff2af155559f99', 'sensen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x34e4e1194f6d2c71019b6eb02f0a629f14aa213b', 'darren.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa03239b8c59df2c79665d8cfbf3a6c1a9b2e5668', 'link.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfc4b7420c1a9b969a7570e052f5f860d3c068926', 'kekeq.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6b8d5f25d86e7a2cde02d3dbb93bf3d58771d2c2', 'ocean.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x863c52d13809e0a05dc205b7cf16604653886258', '0xed.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7ea1eaa27b313d04d359bf3e654fe927376e31bb', 'fatfingererr.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9e3fb8bcd43594f5084a5833abc720dee954e052', 'piggest.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9c86f8c140c8c8c4eef3ddc18711d436ffaac661', 'tedwang147.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc6c68296f0d9694c45977f5842303568d0f8b90c', 'diamond_dz.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xac74fc7db4c034cf8d280ab2256c991bbdc8f1b7', 'mumu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe7fbbc8fccd0062d0928544ee7ef89abdb3b8787', 'mia_qiu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x15de3d7f7180f554421a91e27ee28a542881740d', 'scott.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0db30505c112c376aacd965cca580efe929c78ed', 'claire.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb68f52fe2583b5a568e7e57dc98c69d93821f6e4', 'taosu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x45c4224b8cdc16c2faf83afd541aff840dd3f571', 'trialab.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb5b52d469ac3def7beabf146a204f2efe5010b0b', 'rancher.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa82082380585489b0456e15343c23809bc334709', 'flash.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2b597a38cca6db6e934c44ef0ebd800f015ce973', 'runto.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x018615b3e5f30944dc769ce8882d61b66f31dda3', 'soulcai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x183f09c3ce99c02118c570e03808476b22d63191', '3191.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x77bbfa1e9088b6763fdf18d4655421e81afaacee', 'vera_sir.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa2f87dc5ec6f659dc7a13c8f12663d251e72b698', 'b698.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x45b5e0299401f38d7d18bdde102ae62cdff86429', 'saku.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x49ef9469e3b8160c992bc01e9dc4e710fe3253c9', 'bogon.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x615d174ba9ba4848126426113e617bb97cd0b990', '0xmech.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa86a218cdee25ed38d6e52bcb8f770cf06ab08dc', 'manfred.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfe34816f2455644bfbfde9165a9d493fb9202ce6', 'mowen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2286d9abc64febc22366fd4c608c56d611064892', 'and1ng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xaaca91a07aa078b8002788292c05450c208ee3ce', 'debard.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2cf5ce300fa6e524538b1d7e481ab6fbb270427e', 'hanlili.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0e5f40dc367d757f250b716cd46b6fbde90a1b8b', 'sigs.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8605ad47bd35cafedc05cd35c70b286bfd5fa791', 'leo7283.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb4d9ab4f84a196712dddce620d9d5f2182430c04', 'tonkabean.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3dd4a1f59f3ede92885d3b89a452077961a9c37a', 'icebee321.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x364ae797f7c2564427808dd66470157ba7caf37b', 'gavin0906.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe59a46eb75194cc4411c43bd2e56f049ea169518', 'dorianjhae.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcf326a8369f3da66846a94be77b94b7f07d707fd', 'taxiria.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf1017f7c4fb949cc94df503c80708078061c38f4', 'lkxed.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0be733ac138cf56f4457089ac4f3d23deb70fa69', 'echoec.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1ee993a70bf0bd859ca5ab7a078a8ef31aa1242b', 'hash.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x749e379d6ca9344ee157695ba1dc90bc0ab1daac', 'deonte.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x327b0c7edb1effd6d1325da8d9f2ce92d633c840', 'saber.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x78c0a25ccc21604b3d117100de7c9523f53236c7', 'benlai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x22455ae1ccc04ab4165e338d673eaeabde304b42', 'yydssdyy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa108a44d2a2c65b507ce71bad7ea8fc8a4f6843f', 'zhoujihua.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa55cc646d230f481f80ffefc65b13cb72ed7a8d7', 'liam.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8ed8e47810279b6c9eb35c057bc40d00834956ed', 'v587.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6e7bbd4cb131ea165f0f1f28c426947d0811f721', 'watson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x141dcd04cc497c4d700951944f50f2258491b074', 'utowoo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbd3a320db1ec0207531392a114dd52f918eb30e2', 'jiangjun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x67caf5032a0138628a3d6a2b04ee6318d9afbc24', 'aiyin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf520ace38ecd2683da5f5295841476b22da965e5', 'dinglei.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x36caaf87391031783250206ad80d779619aa758a', 'ycat.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5417c4f05e6ebf09a58a1d6d67a29118e205e2f0', '0x88.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa3422fe244c0eaf84f76b3a29634abcb05e1e0c6', 'frankfika.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa04a738e0034a47072f2d01e4a7dea0bfd6cba6a', 'kybrian.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xae41675d8224a4ba7fc24bcde2c3ba4c3bb4dde4', 'harvey.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x66f75dca1d49bd95b8579d1b16727a81839c987c', 'error404.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xffff1bb3e4778d4822b3a95def4e938c3e118141', 'gurner.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3c6f4ae90fd846e6e10d32ba226096a0f0196a9a', 'vivi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x029642d4960085b7b29894ad10878ee2fb905541', 'hapa.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf52cd2c7c3291e088a027ae368f51ce267d5ffba', 'sophia.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x85fe5b997e38fdac384877e1687b5d4de880c52e', 'yuna.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd12430208987cf5a66ea4a257c34d1dee722b369', 'onetwo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa9d65834adb9aced4f547ce652ca3c5107c4f151', 'pennyx.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x31e292ba33e01c3a9d3467890e2e5c99e1d6562b', 'joann.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4e2b02ed4a3fcdb3cb97cf43f55b3e576ab04db7', 'web3hyde.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6beb239fb74a1c716d970ee78416ca0f979a749b', 'soc_lu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x66f64d021ee1491054b5162569191cd2e4ba8653', 'stone.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8ac08eb8382e058adfbbbaa025ea8ae9678f8d08', 'ergo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5331bde19eb598f2e026cb8c8e928ef82d8ef1fe', 'adfrog.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x26efca99a3a894dcba53a152c9e851960196a0d8', 'laoyang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfb824c36a2b501976bc91f959a15095243888888', 'yase.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x35fd35b113d587a40efc8b2fc7d348df5113b2ea', '888888.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc96c53444c41dfe8400fc3917bf71aaa00d696ca', 'leslie.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x22f24074e40117a109068ff120973b00156dc207', 'rgb++.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x971bc9b53841d589b2fa784daad283b286471a1e', '0xabc.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x59ea66b50668d54364450f2a4d725af86ff640fa', 'cryptokc.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4e49d71fd18c2b8fe981919f233a8b6ee8297067', 'catcash.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x18e9d74f8a2ca8eff3ce49f42a2cf7e7293c67a2', 'moruscap.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x884afe0d1c6dd04f617aad9e6fc7729153b2e660', 'cliffteezy7.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbfc6e7af4fbbc33ce1cd402d68642c465b88d00f', 'huaxiansheng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf7ac63cea00286d65848e9ce4750b9f0fcfea2eb', 'ricy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x14a189ddddb3c80713341c97f194c836b15b8c63', 'larkin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd9bc433fd5bcf0cae561813969c7a69e2200058c', 'kk1314.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2cdbca9766d33d06f203f759ec4f286395889adf', 'darren00.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x91e51f176a34820e6c1f3298847b83a3f4cbcaef', 'xppp.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe5cad2a9da3dde4a353fc3853c8ed68d751c3f59', 'sigmund.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8dab4d84b0ad946f143f1f057acb7bd3a04d4723', 'lane.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb278754d0fd7d39d69a52197a0e53b8d9f0923e9', 'truman.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x35406a6ec3fec114bfc2b6d97da70de0645f8bb4', 'jerome.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb1b8138a03c9b2bfea9cbe9d7442d007fa91fea9', 'tony.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9fbb214b53b7ce372a0aef6c927693d8b6fd9629', 'lens.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x232f1be803e9dff69ec8119366a82b954b1ff8ec', 'fawn.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0238e0ec97128052a020d82f7afb76e12b748ea2', 'basil.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfe042edaed0344b3888fffb8b09dbfef09331291', 'jeffrey.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3e08a265d4a17b1faf35e2dfaca78c536bd0dc2d', 'guys.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4af75a05fc552ca861e599aeb8696d668eb0e075', 'bradley.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x435e87c04f1fc55a074ebdb5bae336ebd222a3e5', 'odin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xadf71973ed86704d06edce48c3ceab738d3604f8', 'dean.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x810384b9023cc68fb7e38f0ba8745179af31c918', 'kotracy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x93a382d9fe207e3ed057bd0b57cba298014c1ae5', 'aikun9527.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x690e91e14be2c7668bf769481362239d1613ddf0', 'papaya.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x504f10249a54421f08b7c4b585b7f11225154747', 'xyzs.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x513554bcb852e0609559dc2dd4120d37895a8ccf', 'kotracy24.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7e983be9bd860d3905cfb90ca2c676993a47a18d', 'shumyoumyou.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4013366e828effe9a1b7f3773af32069a783eaa9', 'binance.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd3c3bf0771e146a0050b42f66e517440ec0103b4', 'unicom.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0e5e7a4f0cd1def303958f40f8697a4ee4181ef3', 'dddd.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb32d45828b0f561f48ffc33cfd9107ba3d764fb8', '18888.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2409f87be5b21bf64cb16094d82a647d63a819da', 'ziscs.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x25680dd1d8c58d95634a139cfc00e0385811e48d', 'niubi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x18b6f61b4b16ccb1715a7a52b89abeee40eb1af9', '6666.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4cdad032a994a6cd628a6529dcd9cabcb870485b', '88888.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x482a80609cff149255c3fcd0fdc9c886d49d4285', '66666.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8e98d77915b5670c3fe96f30389058eb58a8fbc2', 'nihao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf5ddf3c08019b8a2b41cb0a41fcb5a53da6cc011', 'tfboys.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9d6ea530c9274a0c3f6f0626634f8695904d41c2', 'dsfs.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4a2d7f202ef84c23bcb4a446d7b7c75404c24e9a', '5555.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe2721cace0c9055f46529c6bf8aa35361d7de320', '7777.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x23ffe77c28e2a3300d90568ff46a9863e33bc2bc', 'rucky.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdb9712c58a57e0a1eb07912e1acf7aecb0cb2c9b', 'ammi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xee311ab618c846f01c9005e780ff1539b9ed4dc9', '0001.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc103d990c83110e71f77f6ee669d25e44954d0ac', 'llm1888.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4e7d39d59089f8d7c2ad17d61ebd5869ba6b4dfa', 'luckydog.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf7fa46896f6c233bae1b5f5566e3ba9c008a476d', 'wang1666.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4fc76bfdb3f9151ca19895ec41375bf8a087667a', 'liweiying1618.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x60d168257d49fc27c87cab0ea75c34a08f3dce8b', 'worldpeace.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf4ec678f400438bc0818933dd65e8f76f2d52530', 'terraceng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0570f40ec222ec624bfd15a7f18152bb50760500', 'pierce.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc5a63f32a9aa14c629cf856b25ff2d8d5b94aa1b', 'chenyu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa24ab19f515ccc24ec0bb257be21ce984dfa2439', 'dragon.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xeaa76e265b1692ba00bcc2e2584e9312f103f59a', 'terraceng998.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8961c3b806c64fca262407fc164c1c67a7d9445a', '0x0513.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x974d94059a5043229f2e04a16f9bef3b0679132a', 'singapore.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0f129183b97151c5e802bb39f044f1506bda8943', 'seal.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfedfaeccca150a431b7ec07665fdc3d5984bfe64', 'rgb+.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdc03a02610a7ed8975018dee41037f48998b8944', 'superhang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1f8ee3d9b4245cb41ff37930fd3d7317be05eea2', 'wwwtest.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb235e455edab4a2c5b954b9d75d6dbe69ea05fce', 'rgb+++.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6cd70ffeb89373d89712bb5347e34f4680497586', 'gaohongxiang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x77c1712d9eb99b4c630bebbe45f8e952d461aaca', 'fewf.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x61691d1b0a128546eae85458e586d680b3a072cb', 'gerr.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x35f262343e9c10d89f8381ad743d901a1461b369', 'zijiren.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5ed562d506e5bfad38844cdfe3ec51332e934d3c', 'daad.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0f5ad55ae1ceeae83b3a7f350a0b249b6b07e905', 'uniswap.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdd0f068296819bd1bf820df2144694076d0541e8', 'zhumiaomiao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6b26b4d3ded1caf9be76b9eafed89269b5cf3a53', 'notice.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf792f7cc73762dd2ca4de911694486cdee053200', 'sanya.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x726d714b8b6e10b6e25dd7926c92656046d9861d', 'neal.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x13394f5d1dbe125fc4cef62ada690838e1c6aea5', 'baiziliang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9e60cfa5e515272ac5f84e0342ab20dc8dda268d', 'luckyman998.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc197a1d6e5433b0dd53a06f309f8de878fe00729', 'hiimfiendj.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x560a3e60e3baed7f61f715d42caa1011fd0979af', 'hawksmoort.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x420a4e300d74404a3bc337540040713dbb26d3ce', 'joshdevi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xacea48d9d5b13d29ec21ae2053875995e4587325', 'fermodusev.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xecfaec2fea2198021893f849c6871a19f5b53546', 'ghostlanda.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x42a6d36f8c4bd8746926e27a8b56392deac55b80', 'luckybb520.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x336d8bece77c8b87ecdc16c8e912d716eb5d1c21', 'hojasobitp.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1e0b20391631c0f948fccd9a08486da86240657a', 'gloriahidk.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf04931e43cbf070105c373f0958de5f8f1e78fb8', 'drsiliconp.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5e49fcc91d67b0695b8798fc985625b432a57965', 'hergara32.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x494e3ad6dc1e9070be09e054bfb809905f95c9f6', '7687y75.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x19cbd4ecc736f6271e0b72eba784b87c70687e95', 'sky520.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2488df8a782423a65f6afdedbaac19c447e3b2ee', 'ieatpastap.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbba63e36a05a80a88999af66b4d387033b9c7ae2', 'fungiblesd.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x85c2e939d0261d587402e4d29b5e75af0afa4e9f', 'brb543f.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xaafc5e51da8d240cfba6cbd02bbce942eae1002a', 'elronaldos.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6baa255c0eab619676d473355a141494ab21cd18', '7698y97.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xae0d73569aeb1ee1c178445a20c8849f7b33978f', 'dezso.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8c5b2d2f18c5b032d11b09ac1e6490b6856a38ac', 'happykarpo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2c8b93c21d5f9156a939166c1c03df9e68cc4c7b', '4357t8g8.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1f2452e54ba16cc45fefe1dec622b0c1bbb0599c', '6557886.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x665a550a00419195ea95667257e62ee483424535', 'gibi7i8i.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9bc3fd9b726770e1430622e86ee68f0306972a18', 'hernerdezc.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3dbfe3aeeeb64618805fa8af480f62d0a5075d42', '975t54454.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc29113ce5c043a48068acc94fdb1b8482340e0d8', '76757868.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x32433713f7f6333ee6d3b0bc7b5c10f2ccdcca7d', '96rtr768.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3690c9c455f6092b9ecd23c15c28e5443eaeeb04', '65734554.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xba976b8f86bd02cbbb3ff9406b6c9b5a5d34c00d', 'kinze.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8cdcfffb937f73089b5147a38ac40d8cdd738da6', 'blue.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xef0f49c09f502efbb5be9b12e2e391a44653e3fd', 'jazzy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5e87277b3a2e21bd725f43489b25c1c4ae9722c3', 'kejunter.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb5b316fa02da530f6e0b6460026c343117d25ba7', 'bongning621.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0b07f314fd496374f2291249c938b8400d06720d', 'babyboi969696.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf32564e249b695233615aca0dd06e32d67953ea4', '1637.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa6ab19524db770a68950436bd45b81ca8446dc89', '1322.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe392bd6515ba71ee2268eb01d7d907950893ed90', '2863.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xad81586d93d72554a7d768a7654690b5937881ae', '6192.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd6447ec8410623af40786c947136b2135dfafa21', '1122.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa8f4ff15424461ab9a1bcd109b2f5abe1923f87e', '2188.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd0c27ff45193144debc84027c8c2f2eb77ba8898', '1881.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x25e6f67f11f927c1c92a34223bdd11ab4bbfb118', '6688.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x71e284725247a4b6197ca9fd3131a99fca7051cd', '8659.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2b0eb9cb44c717784a4879466fb7fa2d5a214b33', '8168.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x112536c4d8bc18f3004a59b7591b0ce21e2e90a2', '5566.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xeb57e7519fd2b97e35347e105fbc8cca777dd6f6', 'doge.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa6b02c77f6eb61fb06fa04d3356a9dd0088262c9', 'godliiiiin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6bfebc6e2966db98312bebcb44a74d69e3385f34', 'hamburgerk.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x520fee8996bf430e51694a210b831f8c025ef189', 'p5av7x.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x558ed78a8ca9b69d59f0b6193c1f78fb4a8f7b88', '1868.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x562716c194e01669a7ef4233d71646e7d51870f7', 'ctrtsfkmseauk.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x04f4d0337e2567489f033a3cf57cfcbbfaff5d0d', 'tapakipk.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x85e69984411a39a49be68956b2ebfc999f4a1883', 'winitallj.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x852fa6611c03b99c49c854da3228e6548f7bff2f', 'bhagat281g.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x18b0d7e440348f09a8d445d4643889666034e1fe', 'sparkiei.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb0a36934a242497f84f373897bda167497b8ac39', 'monk3yyg.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4a52d5a126c8aa37cbe0842e6a66d0d86ea63dea', 'yeongwonq.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x000d80ac6627e85783df667018c775589b495653', '1989.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3510cce7ed600027095184b0fa4872879847c691', 'aliyumuazr.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xba7be9e10a1c930573c667c9f998c1a4566861ba', 'ryujiiel.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7d12bc1e4ed1f171b21fec19691215de08304616', 'billylioum.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1fd00885df1c24aec7ebe239b9b22b3db039dccb', 'mumu88.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x774e127e05b053edee4b0361cde8abb61c2bb5bc', 'flyaway.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x53cd31587bb9a05cfe5129bed170aec8a397232f', 'yanzhensong.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xecc7200fd4b2b31bfdb6fdbe9e3aeb2e10a01ef4', 'doudou.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xeadb96718bb3b0436d2873ac010802ebc6fee80f', 'danf.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9cbd3d6a36cf2d45442fd35dc160c97f658f8b18', 'zhaohong.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7e1fc2bb7a575d5e641825cae90e27afba1c1603', 'bing.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x68d2067c4c82aa99490f5b4be411ceb71a615622', 'jarvischan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb8280f00ed7fdf2c1add7da632ece521be48f763', 'cherrydada.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6164b87f441765e2e27f14362a7312e18e79293c', 'jackd.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x941d685530d36db8292c17938104d721ba6782e7', 'nero.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe0b897d40311cc72687ab5e0b58746bbfca5dfae', 'lizy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xad8d2354acd5d2b2f56abf5d7ef618d287679e52', 'xiejun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb022fd04b86d2887380067c960637100dbfb82ac', 'golem.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9cade5feae635456332aa7bb9ed77ef8ab2cac41', 'rico.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x73462e3868d0949831b3c7115577408767cbdf38', 'taoist.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8cbf78ae1ab3f6d288bd3ee44c8a74e2a013e378', 'jiaze.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf3c32470824b47732132c92cf11004b9715a14bb', 'binbin.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb9175abb60de2c47c8f1dac88179dda3db369457', 'stsoda.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5d2abdbcc951d39531ec0ae5fc6b438d516d2fe6', 'tristan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x71f80ac94831217629babfa09c5a9e68b3fb07c4', 'markmeng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7373af10bb8c997c44b98f926779d3e1ae455f5b', 'duke.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x233ac30a896c5abb30a21e98c755d8b0cd463dfb', 'soleil.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x045a47cf599aa66c5557fe9eb9c5006ed5298906', 'jl_lelouch.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd47aec165d7469e266f650a049c973ec95c706f6', 'water.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x332772fce634d38cdfc649bee923af52c9b6a2e5', 'treap.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x744a0a58d862a6f7a81d59600207a51d3d40ada2', 'yunashen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x35956416b34236d3f120682bfc9b2b8d4d4544a9', 'toprince.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x798cde64455df6cbbabf2fba57cb7d3c7464bf7a', 'hannah.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x567f0a44a48a378f2e6fa9261c23d6f9b943cee7', 'kelly.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb3c88697d659ff8012872967f46c492d12457242', 'ourai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb7f86bf628b4efcc2ae4d5c36ba4aee340cbdb4e', 'ruth_y_2201.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf4f7d2ed908afd6067d16512aed961e8976db0ae', 'serene.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x60859ed309e5553b39bd66edf1a3c26473da3b6b', 'xiaoyao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa26a6a5d46dbdf4093ccd1782d839a71c150b0c0', 'connie.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2eafe610e240f977aee740473dfc912f6150870b', 'kkek.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd916c2182d2c3bfec7007590e5338a305cd497b4', 'phantom.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x468fd68b81475ac46d00f491bfa52ef8aaa17e97', 'daii.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1996abb5d871cfc408e2e1bd7857e6af2318c2e9', 'yunxiang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc70025f24be879be9258ac41932bae873bf7ff0a', 'quzi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2011862879ee9f98f593c79e5c14bb5325306d75', 'gony.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x358ffef4f02d829d013c46d4a60b7bca3b224063', 'xianyun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x05b45a60256cee0d7759e0ff904e7bc38264e98b', 'xianyun567.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x42db39ed8fb8e672746c8b7db3cc65d4153d680a', 'chenbo0731.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbc9f5417327f989f2e3ed9674d4a3c1d3e526127', 'pianpian.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc8e7e642733e921aee67589309ecf94e0fc07e99', 'mercury.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x65e9cfa11aead9aacdb55daba9a75e506f8769c5', 'liltu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7db5b081b65fb792979e2abbcc852250226ba0e0', 'smartdiscovery.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x1138e94cfa05744c7c00aaf9fb86df9d8b39ef14', 'loxia.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xed43f03fa4162df8012646c05f12b7d0092d3bb2', 'artii.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf19e2d3146c5a97131171e6a57652dfa2014bc57', 'haogg.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2f0544bb9408afe2aeb6144ac93cc5e7d453d847', 'dada.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4f1a79085a36a2e1eb300330734d6a1143c3f3e5', 'zeep.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xebc8f64e864fff699b6a57716f40902e3667c92e', 'alexshi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xab3fa22cf7f2f4f26b7e9b0c36471456e62c6479', 'fosal.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc314d6b97d5476c2d251f51cd7a3718729992104', 'gary.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa9145d47de36c33561663c7c1c82d86a6f7e328d', 'fundian.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x00bd50dda1a8916b98811e4e380546ad6dccd293', 'fugui.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x04ccdaba54a05615dcace5f094315b25e22be629', 'hwen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x30734d5ae63a2f368acfb3d3a15e180169254b8a', 'ping.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x25337fffd30bbb4f28d58b8927895985cb72ed8c', 'hazwl122.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x129331d9caba7acb9bd699581ebdae915ad46035', 'bigfei.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0968a69f08dc1190b797909c96937b8b449659fc', 'qsun.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9fe403723c343697c819c0109a17353b92b1daa6', 'syntonikka.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa8dd786f84f79f61b38ed2a6d0cb99a8be8cae9a', 'aiweb3.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0e6ff2636ae743edcdd1c157b447b591b10f5469', 'chengyao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa1d8518057bb34ec0eee29110b62d97dfccbd41e', 'choeygit.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc0f28572c8fd2c55f0730517cf92310418fbd4b4', 'jy++c.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6df4c584482bf583a252e22b521c7743aa2e6f23', 'tang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x174a5619dbec705681d495979915f11394c5bd13', 'lynn_yang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x81f8843b3f34cc567481c5db17f4238e5d874875', 'devin29.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc5f960ebb9cef1ffa1e399b08fcb77c51510aa63', 'taoshine.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3bba9777a84f7f024c6d2fa3ef1d8ab2128d25a0', 'dido.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5ba75f8a0a01385d75b6fbbfdfa9409c82486141', 'fenghua.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc56b9ad781254a5349c321cefa0a769020c74594', 'pine.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3ed1bcecf597ae3c829de26dd28115c1df0a6b60', 'hongdou.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6af8fa7268f7d4a13c88914c5f5d789eb971c05c', 'zhou.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x623a728aa986b62a0af4888ce5b7f40b48926bcb', 'wangyan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4259d7396a9d82ea0b88230d391d5ee979c49ed8', 'acheng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x955da9462c4960cf9aed82b111c12fb13704e09a', 'shirokoma.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5a0037e6c5cf5ed7ead7d920075f6b558e203ff0', 'justfish.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x84efc8925c4fd657759dac68e4e37701e1793653', 'seven77.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe6714e7ad2460762eba0da6c9dfb47430771928d', 'liu_bef1.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x27bde7e1af100ae0f52926f73f5508137c020df6', 'miao_1922.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7aae4f27f2e724e471a8e4abcbfb62c2d015b906', 'fanka.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x098e4e663d1752afcc9d07fd78c5123f2b0bef1d', 'zephyr0754.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9b79c2e3b8a142c8d131faaaa3df46110191d6f7', 'andy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb65e3faba1277bf5b9de6bebfe584457935384ac', 'bayacat.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7ecccbe9c2b6066ae4789192cee4a512603bbf9c', '0xrory.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe16767f8a662ed3a9789958b4eba7a494944677c', 'flatulist.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4cd9f44ebc11e85586c2422df5dd0d1f0dc58fc7', 'alma.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x7f56d787b7db532fdd4324f7a2b85fa3b3e2c27a', 'kakaroto.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2095d51327621ba664e0e26b170b19108d607fd9', 'kampter.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcaf490176fb075c864db3b159cdfb91dcc86e88d', 'torsten.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4bc4203da39ebea5d6b6b03b8cf9ab7be892011b', 'firelps.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6b13cbc84d149457d08bdd94fd864beba6e79b1a', 'liuer.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9a7cfefa0b22d5f93f16951e79cd537c9c1b76bf', 'leo66.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x90485836d28dd707e8eb46b6d91715e17bd0fa20', 'jessielovely.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe01f4217875ceb96667fce4c890c986691ef6d32', 'zhenyu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x73dba37718e88909135f0b64dc380fb9522d8f47', 'rainbow.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8d14ab0088e4c83f810b74cbd773f19c45f4a146', 'elbon.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf5d474215bd0bc03b11b2d7dafa835d441f73a1e', 'cheng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9ca246434fd2c3cbdea96a94f8e2966e46c4b1c5', 'gauss.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcb0f562549fabee9965eae9a93c51e1242463329', 'sihaihu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdd99971345370591d6d163f244849fa9caf3fad6', 'zennta.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9640995f41054d8251dc3e3568ff0f526a3dec4a', 'xueyuanhuang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x075e539ed17beb139502465499437c5d0a129272', 'tszs1500.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe2b8ac8f74dea76753c3f0a4729f53ce67cd3c75', 'reduo0531.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x709f097a9f06474c53cdaa2715c96c661c8ac640', 'mike.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa4c27285d70a2bdba45e1b0522395a8c59af00a3', 'drizzle.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdc0fb244722383bb12fb9970bc46e60c94fd78f7', 'august.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5bfbaa1051c1dc7fa4a109c9015f131cd86a2017', 'yuri.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa94ea794972aff1d91115ff8732876b780049433', 'fazai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4c9786aa08162ffc838913683eb8fe0acfcb3a2c', 'smurfs626.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x39fdf6f2f57bf028ccb8bdae8bbb2294f2cfdcf6', 'nami.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc09ccc154641467d984f44a7d98a669884c90813', 'nico.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x99ffaa95fdb0a1979e9631882dca58f6fcdd41f0', 'qinchuan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x005fcbc873a721bf6c2c3e1c12348220f2d7455e', 'zoro.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfe93f387d07c1242e969b085ceefe2fface41d2a', 'tiger.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x674f64a390319393911a31573c301695cf7c93dd', 'lubi.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf95bdaaa2391ecedfd5cbb32b51c86360b0654f2', 'zedz.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x59b62cfef084740ebbd00bc5c16c3ae0e730518b', '1284475!djhf.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa9de694693ae1fb9263f37249b4989d4331a3c36', 'flyingpig.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x419ded46235fad0c371b652f3d85a82234d21f60', 'jerryliu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x728916627585e579a7fd33ebc2232671cc5fd651', 'fangfang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x016646923ecfc2ea7cba5ab7518b7eefe2fba73b', 'luckymario.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x40584cca567b4a57ee85cf526021edff0a678bb0', 'chan_123.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfaf219488c82f752a4e8f77ee8809d95d2a2caea', 'roger.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2052bfa53ad0c294d6873f9056f6a81679e35a89', 'helen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5e641ea2e28603edae61d7a6584f78c03e1f8a77', 'sihailieguo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb7789200c2ee8d0d8221b03c6c3655fe8611e004', '+waste*sea+_-.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0b39abc301236be9dbe9d2c78a74ea9e9cf7e6d1', 'tt.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4be79c2d9b700c4863760e4b95ef7cd94565516e', 'mxiand.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x4863e40256f477ca540702f608beea67e0f1f896', 'loveforgod.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x50fd442b90343aa7d484c96588e765e4e80c1ccf', 'alfred.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x65a3d5484751120f727a190b56dbceec6a0c7a7c', 'mentalization.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdf3f6f097f648816c8fb87130e7fb59554d35250', 'chen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x87a53df3c68864b08b2fb4f82475b5c6acd9987f', 'anson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x540fcbaea29ce96525451130996c44f113d9b605', 'ccwinter.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x045be945126bd5b361725c25d0cdbbd074944796', 'xiaohan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x53df83f516d9f998f43f54287c314c2698fb42c5', 'kiko.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x70ebd8f1ca89ff78f3f88f53943bb30a0e32f44c', 'carollyn.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x38a5fbfb3436fbe93dfbf09efda594c2ccaad6b4', 'zio_31415926.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0cecdbf823ab1ac18fb7e01ffe853fe23721b22c', 'stan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcade70a43ce1f170d717aa139aae3eb581af5783', 'lenville.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x967fd257d50b810cc8b47c7dbdad7c96316e0067', 'liu1997.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xffc4dccb9e3ec9f31a5540e27a68f0bc06f5f572', 'xiaobai.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc331d54c508bd018296f8f0056275cef75b59b8a', 'rickychen7.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5e6cce07a609d7550ffd39beea0d8b2eef28fcd2', 'cyaneumluna.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x70382d00c9c565bb4eb05edb3eaec0f00412c6de', 'emmalove.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xc84b4ca8011310bc4bf44f2d1dbd881ee7f44aaa', 'qyuhui.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa844937990d6f00badacb4a278e6233c1db47310', 'soulc.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfd1a416ab465673e6a185559b8c111b8a47d631c', 'hyan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x242130d48d6173d94f519f8371708d31f9c710c7', 'oscar.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x054b51dde083307aff1888de90c4278ae50da137', 'littlemodesty.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3f37aadf55654616a3a2c63d5093b787536c52cb', 'ravenw.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x831e8e37ae66636193d070d37bc601d1f30fe0b9', 'bright.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x65ebb979c7a0ea83b37aa02ab0ed4c9ca6b0cd29', 'bettyliuliu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x832e6c85b961775cdb6b4196f5d43f3f740bc2e6', 'sylvie.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdcfebe7d1b4d1c0959a55d1b7b9a945a7552486f', 'neverink.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xe23ddc2f7031529dd6ea3d48c06a6b6be0dc02de', 'wallet.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf4f1b1146a08fb0d1fca4807843d1dc5ae2872cb', 'azhan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2679d71b87a92e9dd0632ce285f02ad4a608dd24', 'yunshen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x95dd90fc088067dc161ac63ec65bf4542a219ea4', 'drewdy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcf2cdd019cf59ff94692c457753998c0457b85db', 'drewy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x20a738ab28957e25f62e20902ac4ec8d4e5c6527', 'zephyrelis.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xbacbb1bf7a51e8a4d3fca21ae1f6755e2e1b8449', 'ansir.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa4d945f97c088a31df42450e4b3166d75d5ecbb8', 'redcrap.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd687260f1d9249466efa009a7ce14edb60a81f98', 'theo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb55e9dabe4b64e5cb217d28dadd3a4f815774e39', 'kkkkm.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2fa196027d337a4e5ff1dfafb9881a2726ab9b02', 'soundrain.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2bd140bb0d04263eaa611006a8681af55dc8e051', 'sanji.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xddf920f9ea110e1a8030c4015d17291c7befc079', 'tdf514.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd46ebfb04d3bbb3b7bafa4e40cd39feb4cdec278', 'roson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb1df7c1713292425646b23e00eea1c31090e624b', 'bella.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb14ff8b229fb46263d33ab8029e4de3ad98c49cd', 'starrydesert.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x794c99a21a5d99e3a43820da3f3f17dca427cf31', 'along.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9386d783483efa2a864050279e8fd952fa4cf1fc', 'ding.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x632a9e30136f710a2973a0286224e206778bff07', 'flora.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x36ffae59d083fc67e17645330d06d8c3ce3e38c6', 'lily.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2456d6ab70a060e58851a443e44e08d1e2453aa5', 'pangugu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x77f41eb2016e06e1cd760ba67da8f599f98202d2', 'jessica.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xcedb14b5cb7065744e335fbb5be767ef3207dccd', 'huihan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfb38f426c497d60c3976473245cce04d185f72f7', 'sunburst.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x431b4ca18e269fc7e1f5af49b9f4e2af683f6207', 'd-lyw.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2c50a0813297deda62df66f0996383908c2b4faf', 'thalesliu.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf6fc63c1cc6e2cddf871a4193d676565116454b6', 'zys11011.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x72ea2e95a78da955fc85b35576d5e9007c9b4db5', 'sifotd.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xaf25777f60eb6aa69538fbb91f444320a0fe4a4e', 'silly.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x67b7824628d741d6979009096c12f0abfd92335a', 'river.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x185d4f23944fb10b16505b2ac8f2e94bd4bd5874', 'cleo.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb32886b08c45e9aa9bb229d189f609e3f837e8d4', 'xiaomage.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x858e2492c8a2cce519d30031fb09da89c4043a1f', 'samantha.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3939a8d0c549295862dbfdb1e1cb8bfb4a118627', 'john.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x132d7968d98ea1a4a968c36b4d13af0293e19477', 'bigjohnson.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3ea06f44b81cf15dfe403767ab9ade99f8ddda68', '0616.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x833c72d4d6f7abbc7dd9ff44237c6f3d2f14983a', 'ryan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xa920be1bac26025c303285cdae8d99a90d91f459', 'peggy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x729f30e5b893d2a4d84f4e1fc0c72cf493c4669c', 'jingqixiansheng.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5d9270df3787516b9d6a69036f63e15dc864a6cb', 'honalee.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xacbb6c74410e962b62bc91751fb0fd808bb8f54c', 'lzy2006boy.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb02c5eb954b95aae35cc32337cae4a6924c727da', 'grit.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x88233c3478c569e564442edb9287c67f6a24bd27', 'fang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x016c6ec0f9730249be31867cc8bc0fa6cb00e7d7', 'fc_huang.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x0cd30b39ea15f1455e3371cb1813d0f51a57ae27', 'baisen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x909a8eca017f6d2f7a22b893ac2d3353af45a4ae', 'smiley.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x05bfcbd6936d5dbb6d7e43a0f332dda741847d7e', 'anony_ymous.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x17f839b1cff0c47ddcaf37d78efa6cc124112232', 'hebye_.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5816003626ab51e7823a2aab5fe47e1eb5453810', 'oslob.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x36ff32e01f5229b52fea4b598560b81c96ed9685', 'muye.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x52c5cd96e45c739f0648f5bb42622e0bbd3e7e60', '7888.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8462f831a99f4187296b81772146a3e710dc3506', 'lightningchris.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x975d01a32dc3ad60a0edb0f8f8ddd25d3d60c63f', 'samcc.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x6001fd2b99e550cfe32f1dcfc8045ee565fb15bf', 'fofolicia.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf230619aad1fbbd483705b5663e4adaa92b4bd9d', 'wdx13.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xb633c52435e8269b85d242761a79544d6522762d', '1995.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf128efbb3b96adb1844154d6f61499db01712771', 'shulan.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xf0e7bf74ffa098dc04fb2068be91c554a16c7646', 'burt.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd904ec4d0de7df960916bf3a05773bd49a6cc596', 'galaxia.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x8ff59c2c973c424638129caf8a91ba3b15c6386b', 'zhiyuan65261101.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xfe8496cd25d7e40ce19c58d21a64747796478cd2', 'zedric.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xed8f194f0a116575316e3f58febd40ecf6dcec13', 'shell.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd41a72faeba8837e25a94935e47742ee92a9cd13', 'jiayao.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2ae6b8d78a5c8ccd04c3eae2e2c9a4f878919b25', 'aimi_.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x5515d230ae406f529e93eaec469cf2d69c7e8ade', 'lin2025.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xae9ceda940747c36e912c611c81238ceefd7f267', 'ynjzw.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x2e6a47e7dce425904e09f1c1446dd01be4d0e47d', 'javen.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xd26d4ef9e4351e264f21f58aaecfc4a31cc11946', 'wdoug.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x9c399a051e4c4e2b2ab46a833525cb045435fdec', 'mark.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xdbb63ba78b7f336bda8d65f771fb3574221a4d6a', 'rojazz.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0xaa28b9a697dc4396d287d1583e587b9660d596a7', 'geekcode.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x73be6a24dc296f61f03ac8f4092c3cdf923fc76f', 'ja55per.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x3902b4cdbe43b2c5b425f43d8356b37f3fc3c01f', 'asher.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();


INSERT INTO sns_chain_registry (wallet, sns_name, claimed, synced_at, created_at, updated_at)
VALUES ('0x280dfc29505e51761be17b2a958874a47d54606a', 'iven666.seedao', false, NOW(), NOW(), NOW())
ON CONFLICT (wallet) DO UPDATE SET
  sns_name = EXCLUDED.sns_name,
  synced_at = EXCLUDED.synced_at,
  updated_at = NOW();

COMMIT;
