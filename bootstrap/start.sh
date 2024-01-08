#!/bin/sh

BIN=/usr/local/bin/monetachaind
HOME=/root/monetachain
NODE_NAME=validator
KEYRING_SETTINGS="--keyring-backend file --keyring-dir $HOME"

# $BIN export
# exit

if [ ! -f "$HOME/.monetachain/config/genesis.json" ]; then
  rm -R $HOME/.monetachain
  rm -R $HOME/keyring-file    

  $BIN init validator
  cp /usr/local/bootstrap/genesis/genesis.json $HOME/.monetachain/config/genesis.json

  $BIN genesis add-genesis-account moneta1qx4n4n6t4f2fs8npcxse8r5k5hzv67mhrrpce5 1000000micromoneta
  $BIN genesis add-genesis-account moneta1q848whv2upgtewctqwnerqcewrklx07v5420tf 999968932900micromoneta
  $BIN genesis add-genesis-account moneta1qfyan2uks3xj4hkwcs85en89lkuncnthkjglhx 6000000micromoneta
  $BIN genesis add-genesis-account moneta1qk2f0fklh90dv5quy0qky887amcyryvyc7ga5v 2077105micromoneta
  $BIN genesis add-genesis-account moneta1pd07mezyak88sj0u5n6979rakwrfzl0x443wvq 4131968220micromoneta
  $BIN genesis add-genesis-account moneta1pjpm6dp3q4nywlhe2suxh0mnf2nzwv5ezt5d80 130000000micromoneta
  $BIN genesis add-genesis-account moneta1z9pf67drhfv0r2nulneu7vdz9y7fvsmcg5x8v4 7000000micromoneta
  $BIN genesis add-genesis-account moneta1zfknm3fr769dk359ett967u475kenvykyutgw4 8000000micromoneta
  $BIN genesis add-genesis-account moneta1z3edmnsjnee896n8czeewpcx044h532s62vnga 1000000micromoneta
  $BIN genesis add-genesis-account moneta1zh43gs8jwlu3gnd3gek49x6y6n7urfyfjqmj2q 23351889567micromoneta
  $BIN genesis add-genesis-account moneta1z7376w6q4dn4ela5uvlgmju2053vz3pg8y62a6 1000000micromoneta
  $BIN genesis add-genesis-account moneta1rzux5cwmeqqpfn6ecw509rx9pskd07v52gg3nq 307674522micromoneta
  $BIN genesis add-genesis-account moneta1rxa20f9wq0ees78pl2xmhqdjrxxqg9d7tut7nl 15000000micromoneta
  $BIN genesis add-genesis-account moneta1rvyqrwfghypp5thjfkwppa5mk95nkm0enjzcxw 4929278256micromoneta
  $BIN genesis add-genesis-account moneta1r4rfvqyjj0nycretax744krvrwlh404lp30083 9997847280335244micromoneta
  $BIN genesis add-genesis-account moneta1r4gq3k7v9nhew37enta5a6cv6k4wh4gwyqeacz 1049976699micromoneta
  $BIN genesis add-genesis-account moneta1rhavpmwwj4qegm8j2selsedanjqc548j5jnumv 499000000micromoneta
  $BIN genesis add-genesis-account moneta1y4rfy2pppsv6spg7erus2hwmhhgm6vkrxca9m0 5092258485micromoneta
  $BIN genesis add-genesis-account moneta1y4wr7ehwd22sd4t7dkwyzx8lt84tu43e5nwnem 1107539056micromoneta
  $BIN genesis add-genesis-account moneta1ykydttjjad2598fu9he4hj5vq8rpjxzhnyzwpf 58522652539micromoneta
  $BIN genesis add-genesis-account moneta1y6mtm6cjlmvs27crye9pc4js647yr950q5lehx 5000000micromoneta
  $BIN genesis add-genesis-account moneta19zrus0y3adn049fdr3ppju54rhe6657pa2zt7n 27398540536micromoneta
  $BIN genesis add-genesis-account moneta19w3rp0zzmzkj8kfyadpdxd5r3yqfwc6xjge9zj 22000000micromoneta
  $BIN genesis add-genesis-account moneta19susme2k4r66la7h8fzvl9af2lx2xhnp8wtly5 26000000micromoneta
  $BIN genesis add-genesis-account moneta19jwwhswftpn67tjkkl4gz9njpjqzhztygtcwmh 3000000micromoneta
  $BIN genesis add-genesis-account moneta19uxujvztrv2f9t0lm7ghxpx36k66d5fpfa0wr8 10000000micromoneta
  $BIN genesis add-genesis-account moneta19lng9f69sxh2r4h2tmhu3a2zjl96hw52afl6se 39000000micromoneta
  $BIN genesis add-genesis-account moneta1x8qe7f6falluek3uw66s2jsw5egp74es30c9wa 59479596micromoneta
  $BIN genesis add-genesis-account moneta1x843xhwv2lgsznl4f0e3htur9p672nstz2lwrl 8000000micromoneta
  $BIN genesis add-genesis-account moneta1xfnt930uz3r3ylr99trhudyllqsqt7r2nmux4e 16000000micromoneta
  $BIN genesis add-genesis-account moneta1xj8w437lv3rqxt5vd73grwkvhyrrgfwdq6ql80 499988904micromoneta
  $BIN genesis add-genesis-account moneta1xjasc9pz750ykecmjxv458z48hwg5yf90jagy4 10000000micromoneta
  $BIN genesis add-genesis-account moneta1x7xa6t7mam7a926ya9pjnwdz7p4mdv6e7lpv5w 24000000micromoneta
  $BIN genesis add-genesis-account moneta18qmvn75j90jj9jwak9k70dtehw6672frfyc7ge 3000886918micromoneta
  $BIN genesis add-genesis-account moneta18q7dhxdam3qfdsv4l2mfqwle5qeg0nnear7a57 5000000micromoneta
  $BIN genesis add-genesis-account moneta180kcv6hx8tnz04jkvufjxpzz9glzhvelyqpavh 434855243micromoneta
  $BIN genesis add-genesis-account moneta1847k3sl5tve6e73jd7jxz6at5ercjjkskmj5np 97949000micromoneta
  $BIN genesis add-genesis-account moneta18epacq30x3sqj9r2tm6dgeazdnnemv8dgkdj9f 110534192micromoneta
  $BIN genesis add-genesis-account moneta18afdwasflv6crs7974nrmsyjy7e7ulmy3nfmyu 12000000micromoneta
  $BIN genesis add-genesis-account moneta1gyac6693yexyjzttxnr6e6qt7mfgly8fc3asf2 0micromoneta
  $BIN genesis add-genesis-account moneta1gsw9c8j6vez9sljrsh7kwvy4repfenhp0mqdzk 4000000micromoneta
  $BIN genesis add-genesis-account moneta1ghntf2sala386vd5yd9qr07klg2jz3nfhgehmz 9000000micromoneta
  $BIN genesis add-genesis-account moneta1glr5qyg24hclku4hwy0330mvndr5npfmpkch7p 41000000micromoneta
  $BIN genesis add-genesis-account moneta1fq72kegr7spdqe53dj7shfuk9vpvsvj5sn6cf6 5000000micromoneta
  $BIN genesis add-genesis-account moneta1fzm38dse0y2s6ul540am9ajakvgc2pjlf97hjw 5742242micromoneta
  $BIN genesis add-genesis-account moneta1fyuttznflxzzp678c99fycshp4cqdm275g7e5a 50000000micromoneta
  $BIN genesis add-genesis-account moneta1fv79m7m00mr2d0jm4chxx57fqka4sqtrh3qxra 2000000micromoneta
  $BIN genesis add-genesis-account moneta1fj6h4slwh7u7c6wdlz2d8jnyhrujv97leg4jrh 23182294447micromoneta
  $BIN genesis add-genesis-account moneta1fnsuzalnqyr5854dp6xdf3s2kk59qjnuatk35v 10922895micromoneta
  $BIN genesis add-genesis-account moneta1290n9npnzykmsxj2d35vh0r2mj4y36u2xqpcsf 20000000micromoneta
  $BIN genesis add-genesis-account moneta12vvykwqr3t43twpz8p75vxx6rkuch7u30cjufs 60990178846micromoneta
  $BIN genesis add-genesis-account moneta12vaf8z9n8w9at3rccdzzdzqrs0hya9vf64e2e8 17000000micromoneta
  $BIN genesis add-genesis-account moneta1232gwzqxvrzt966x4r3n96sag7nq762zfcwkjp 4000000micromoneta
  $BIN genesis add-genesis-account moneta12jz3gmkc0r6kvepzdxmrlpctp5udzf3lq0hmm2 117563097micromoneta
  $BIN genesis add-genesis-account moneta12j3hal520rv0mxar6x7xy06l004kzvpql77dme 24936675476micromoneta
  $BIN genesis add-genesis-account moneta12cdjc2wztlgp36e05zazzwqct2f03cj7acwutw 2000000micromoneta
  $BIN genesis add-genesis-account moneta1tsnnrxs5qf27ntsg3zsm7nc5t9sphg2mp2ftpc 97115136178micromoneta
  $BIN genesis add-genesis-account moneta1ts72qcfudhtapjwmytacfj7v2s0yj4dzmakfxp 16000000micromoneta
  $BIN genesis add-genesis-account moneta1t62un2vpqktrttapf3p05xrgq83d6tn9aaa29u 104000000micromoneta
  $BIN genesis add-genesis-account moneta1t6hne7nzg5g3d0rq24z67fwfcztzn3s7cue9gs 7402927113micromoneta
  $BIN genesis add-genesis-account moneta1tl7fqpeg59pfad2hkctu56rkl0lnea35x80gtr 9500354micromoneta
  $BIN genesis add-genesis-account moneta1v9l0vdnhj9nt9hzqu8cnmxkelagg8yme7nrpd8 15214307233micromoneta
  $BIN genesis add-genesis-account moneta1vge0wjzmtu5zcys8vsj2gajl83ruz0nwfxpm52 234115915620micromoneta
  $BIN genesis add-genesis-account moneta1vjr6cx6cxdjlhmsqng8llnqwm9u2mz20ytc3dv 6000000micromoneta
  $BIN genesis add-genesis-account moneta1v4hy8g4sr99h7r4uyvnv3p5atq0vj266pp386n 12660896027micromoneta
  $BIN genesis add-genesis-account moneta1vhekckmejs9tgvamnkhxdzyhzw6hepkemwxq3r 14000000micromoneta
  $BIN genesis add-genesis-account moneta1v6mf0ajs7ggjlqngfjhtdhf92rlsssj80kcafs 150000000micromoneta
  $BIN genesis add-genesis-account moneta1va2aqg699elu7lg6pk33tyruxjsfsdd5mwecjy 15428704586micromoneta
  $BIN genesis add-genesis-account moneta1dvts9380le9a03x5zpa3rcw00xqajm9wxzgjre 87997177547micromoneta
  $BIN genesis add-genesis-account moneta1d0yauuq8a37jzkah3q6elxd5gqwsf6lhexqlly 27000000micromoneta
  $BIN genesis add-genesis-account moneta1djn8v3eu4hg3wnkuprxxx7mrud9euc327q4a6l 1000000micromoneta
  $BIN genesis add-genesis-account moneta1dn92phutmq4erqze3g87lqrz4lkp20dd4pxw70 0micromoneta
  $BIN genesis add-genesis-account moneta1devwck8mcfmx4lnn5zmlc7r8fckdr0v64rhxk2 1000000micromoneta
  $BIN genesis add-genesis-account moneta1d62xy5tsshza5cd3cckwc9d8rxvt02w4yxe9ku 6000000micromoneta
  $BIN genesis add-genesis-account moneta1d7nu2w50mpjzm2cwgwha73pr0tmffnvaptwap9 5000000micromoneta
  $BIN genesis add-genesis-account moneta1wp5v0cgg4ejgntj2hlpnycjp7m50ytskpnz2wg 1000000micromoneta
  $BIN genesis add-genesis-account moneta1wtsrycp2cfswj6er5zzkge698fjc3haew0ucll 5000000micromoneta
  $BIN genesis add-genesis-account moneta1wv3ujk9v5uwu9ft0jt8hne8t75v7n5wvjunjfp 25000000micromoneta
  $BIN genesis add-genesis-account moneta1w4y4zkqg8y68fvhsegy3yvflts2dledpq4nd5q 998000000micromoneta
  $BIN genesis add-genesis-account moneta1wksmernglm9d2lu9wsxnfssnspx7gfjyjdr5u3 77000000micromoneta
  $BIN genesis add-genesis-account moneta1w6gjt2mrwrcs562ym84rzsm9a9tl9yqmxgxp6v 107996534840micromoneta
  $BIN genesis add-genesis-account moneta10qcz69s9chsf0gqdraxw803duv0cpdguqhm6vv 998000000micromoneta
  $BIN genesis add-genesis-account moneta10frau2ggw8vcsgtukqy3lj03vjajzm0e0eq4eh 35000000micromoneta
  $BIN genesis add-genesis-account moneta10vh5jh9yc9q6jf7s24mfn29t2xalx2dvhxlplc 1000000micromoneta
  $BIN genesis add-genesis-account moneta10nwuz54lf5gdwl2hxcmmxsdru0d3mf54yy7hmy 8000000micromoneta
  $BIN genesis add-genesis-account moneta107ewpjaan6vzgagumhnstjklmkn0752n4hev0n 48000000micromoneta
  $BIN genesis add-genesis-account moneta10ljvmxcgqfqssg8rt9pmvg9er3sca3a6wlvq9v 55000000micromoneta
  $BIN genesis add-genesis-account moneta1sqy0mt94wrw0vezdzdjwsmq4fn3wfqt0n8f2ed 1205529422micromoneta
  $BIN genesis add-genesis-account moneta1s85yx85dn570dwxefp9g4c45srcpzuh69fj2p4 298993342micromoneta
  $BIN genesis add-genesis-account moneta1sd0vma6e87gt6d8zrmz7ruvgf3v2e06yew3mfa 5000000micromoneta
  $BIN genesis add-genesis-account moneta1sk6rfj6vn5ypm9ax3hkeuyfxhuxn308wrsczpn 1428884micromoneta
  $BIN genesis add-genesis-account moneta1shqe72psuhpe962ph5cpn9zgg0xgtja3cnzyvw 9422557micromoneta
  $BIN genesis add-genesis-account moneta1sanqpdsn6ygg8z26cvxvaehpt2ecrltdygr0zf 10000111416micromoneta
  $BIN genesis add-genesis-account moneta13vmg8fs4ye8m0fvyzzmal98fa9nfajjrtet6u4 49794083micromoneta
  $BIN genesis add-genesis-account moneta13ne7k9zazel5u47zcuyr9n3kf3t6zyqyjdshms 31784496micromoneta
  $BIN genesis add-genesis-account moneta1jfpphc4k8d9y6k0m3c645hggdex9gh6nw5zk7u 3000000micromoneta
  $BIN genesis add-genesis-account moneta1jwzk8sqm3d87zql9de75m9x8jrj2x5c6z49kw6 21000000micromoneta
  $BIN genesis add-genesis-account moneta1jnqmk8yehl7tt8823lmkvc89ktyvsae59n9sxv 43176189851micromoneta
  $BIN genesis add-genesis-account moneta1nvt7cdxspd4vehu7ryugsp3m8fp7uvrhppnmvj 6000000micromoneta
  $BIN genesis add-genesis-account moneta1njqhtwcj74mrg4jt3ehlsdag75epypxe96u9s9 5000000micromoneta
  $BIN genesis add-genesis-account moneta15qw6qmha0wq33gc4my2x6u2dakznjhvtj2crrp 39627277712micromoneta
  $BIN genesis add-genesis-account moneta152q35ej2uqekqsglet447e6dzmcescz5u8pdz8 2595996388micromoneta
  $BIN genesis add-genesis-account moneta153ykxewq3j8kh2zzhynf6trekgs7vfyvwlcx4h 7169616micromoneta
  $BIN genesis add-genesis-account moneta15eg4emed4e9a2uan7fjylmawh9zfqsac8sa5uq 3000000micromoneta
  $BIN genesis add-genesis-account moneta142hmx3gqyd574lpzwjf7sfp26t4dsxax6wcxxu 16000000micromoneta
  $BIN genesis add-genesis-account moneta145ce8tlellas35r5dhsycyg99fg4gg47asnxuj 4882193micromoneta
  $BIN genesis add-genesis-account moneta1k2dzydrpdaatw7ljgvlf0kslnscu5z37l5j3v9 25000000micromoneta
  $BIN genesis add-genesis-account moneta1k4vy82jxwawsek0aax37ezdylhkxmvt0crtc4t 2000000micromoneta
  $BIN genesis add-genesis-account moneta1k4llmgwsw5wmgx3maeslm5jh73dmurnucpqld0 3374859304micromoneta
  $BIN genesis add-genesis-account moneta1hz9kkmcpgwjrew9safr7a8969s4cs22et4r6vf 4747319710micromoneta
  $BIN genesis add-genesis-account moneta1hkwu5289ue8h766pz78ujhgav2h6r6kraz283l 13000000micromoneta
  $BIN genesis add-genesis-account moneta1hhma2dane4r4q9u6rc37l0duaj4z8g9uequs0y 500000000micromoneta
  $BIN genesis add-genesis-account moneta1cvucr4j9d2ea6a243lwjqnauyn0qqu0cqcrtrz 109327903micromoneta
  $BIN genesis add-genesis-account moneta1ckpjulyw8uwh5nxagj5gdqv9vyevkg2mx002p2 1000000micromoneta
  $BIN genesis add-genesis-account moneta1ejr3z2xcnn9jhxywqvqu97vcgpunfcs44ymzck 21000000micromoneta
  $BIN genesis add-genesis-account moneta1ek8dqps2u36utgrnerlmarsrr5mjzkndxy56uh 159000000micromoneta
  $BIN genesis add-genesis-account moneta1ec50ge7fcv5p62af9pv9twvth52dxvqsd6udk3 97000000micromoneta
  $BIN genesis add-genesis-account moneta1emu2uejz3qf7qtpkekrkkux2nwj2t49c9ukrjn 67100micromoneta
  $BIN genesis add-genesis-account moneta1eupyv6sszgglsycwyawfvncm22zudu4ypkr4m4 29000000micromoneta
  $BIN genesis add-genesis-account moneta16pffjfjs9338r48wa7y2wgz4ztg28vgs3xl0l2 31000000micromoneta
  $BIN genesis add-genesis-account moneta16pss0z77ly9knpvgem0ne4u8e20ypqkp650g8e 15684401011micromoneta
  $BIN genesis add-genesis-account moneta16y8a24zejjjw5unyyt99p3d23wajlej569kgw2 10059776753micromoneta
  $BIN genesis add-genesis-account moneta16trk3pn672wydkufz8djvq6g3r4m7zl3skz6tk 8000000micromoneta
  $BIN genesis add-genesis-account moneta164de0kzpdztwpsdl3tca0d6y5wkzwz05kk7mpa 10913714317micromoneta
  $BIN genesis add-genesis-account moneta1m8x25sn8lqa3n8w7r0t5n940k9effh4yxd2ztd 13000000micromoneta
  $BIN genesis add-genesis-account moneta1m0f26hagvy6p6kqqw6vge0r93ckf8pz2dgdtlz 536163micromoneta
  $BIN genesis add-genesis-account moneta1mehyr99g7ex5yy44gjde73nwvsacr77k9wvvrm 9000000micromoneta
  $BIN genesis add-genesis-account moneta1uqfmfxatq7d3hwcvftgv6j0ydgqeuj7uhj5wjl 1276436014micromoneta
  $BIN genesis add-genesis-account moneta1uptemt8fcmwl7utpf8ks9lf7zvkv8fynkjk9ke 1000000micromoneta
  $BIN genesis add-genesis-account moneta1upkapc0qp7cmz97t2mugpe7uyzdjn7lv4s66dg 26186839183micromoneta
  $BIN genesis add-genesis-account moneta1uxv99fak8hp9sag80d2mw97d6jx0w8c68uduf9 7000000micromoneta
  $BIN genesis add-genesis-account moneta1ufjxrrga6f22ajemd5pzg54gy3hyglg87a8xjn 5117815969micromoneta
  $BIN genesis add-genesis-account moneta1utxvwtsu265evmluak8xyf43fcj0pssvtmwxcl 370465808micromoneta
  $BIN genesis add-genesis-account moneta1uvxrxsvl9t4m8jhxaszjf4sn3c0y0yfdg320ch 6micromoneta
  $BIN genesis add-genesis-account moneta1uv5h60cjuacadsf43cfvpn6ptd4mgm4pfv6p4k 6512158micromoneta
  $BIN genesis add-genesis-account moneta1unxs8cj9329fgv0se2tvzxvky2rj6tgz6vctpt 51000000micromoneta
  $BIN genesis add-genesis-account moneta1ul4pgyv6lnpx8f4xr82n5mt5stekzs5u8v3008 4000000micromoneta
  $BIN genesis add-genesis-account moneta1a9pnza4m5tlwz2lhfxr5tg82shflp0mwgmfs73 85534329micromoneta
  $BIN genesis add-genesis-account moneta1adtjjxmzct8um3338xukr0draufxnxkv8h8jt7 1000000micromoneta
  $BIN genesis add-genesis-account moneta1a3ura3zz5gmmdjhp5rx64n3gt9alucc2mnrut8 15000000micromoneta
  $BIN genesis add-genesis-account moneta1aeay2v7sa2q7mstnrn5yh9kx84jr036pkprqvk 13892871micromoneta
  $BIN genesis add-genesis-account moneta17qe09y32keqlxq4a2z9d69g6ng7nfcda02y8ns 116164175489micromoneta
  $BIN genesis add-genesis-account moneta17yurwj56r3qtva46ujh55xmf92le5jhxc4ln5l 26146054141micromoneta
  $BIN genesis add-genesis-account moneta17v92vpslsg7fdteq3rcnvjhdadqsma8dajzgfq 15000000micromoneta
  $BIN genesis add-genesis-account moneta1737lduj299wfn75yzufqp7sjdwn2muxw70wh2q 16645427micromoneta
  $BIN genesis add-genesis-account moneta17mj6dxmcv8qfx9904m2yma9n050v36ge76y8gc 5000000micromoneta
  $BIN genesis add-genesis-account moneta17at5wvt8xntjy4csrkmk56hzj3eqyhl0ev27gp 7000000micromoneta
  $BIN genesis add-genesis-account moneta1lz86zcawfxma82wly8mdgwe40mwepgzxj4ep4v 13349752901micromoneta
  $BIN genesis add-genesis-account moneta1ld34hx4xjj43kwnn336nwhcfka2yqje4ky8aec 110000000micromoneta
  $BIN genesis add-genesis-account moneta1ldmant5slnaqgvnqln7wgygunahwkqnpsq80lh 8000000micromoneta
  $BIN genesis add-genesis-account moneta1ls2z6t63h0g7pzmc50je8r4gppd8r7zc3udlwt 36000000micromoneta
  $BIN genesis add-genesis-account moneta1l4tpdn0kfhwsajk9j2ecudq35mndue4ssv4ukv 1000000micromoneta
  $BIN genesis add-genesis-account moneta1lmmcl4raefd7z5d6p0pernlypq5zxs9ut3y78f 8000000micromoneta
  $BIN genesis add-genesis-account moneta1lu4j0v32nhw4gw2cg8qrl82e3wsuplze54s7qy 17000000micromoneta

  # if [ -n "$MNEMONIC" ]; then
    printf "$MNEMONIC\n$PASSPHRASE\n$PASSPHRASE" | $BIN keys add $NODE_NAME $KEYRING_SETTINGS --coin-type 707 --recover
  # else
  #   $BIN keys add $NODE_NAME $KEYRING_SETTINGS
  # fi

  $BIN genesis gentx $NODE_NAME 100000000000000micromoneta --gas 1000000 --gas-prices 0.1micromoneta $KEYRING_SETTINGS
  $BIN genesis collect-gentxs

  sed -i -e 's/minimum-gas-prices = ""/minimum-gas-prices = "1micromoneta"/' $HOME/.monetachain/config/app.toml
  sed -i -e 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/' $HOME/.monetachain/config/config.toml
fi

$BIN $@