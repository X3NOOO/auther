# auther

```ascii
 .==_==.   dBBBBBb    dBP dB dBBBBBB dBP dB dBBB dBBBBBb 
:|&&&&&|:  BB                                   dBP      
|&&' '&&|  dBP BB  dBP dBP   dBP  dBBBBB dBBP   dBBBBK   
'\&&.&&/'  dBP  BB dBP_dBP   dBP  dBP dB dBP    dBP  BB  
  ",&,"    dBBBBBB dBBBBBP   dBP  dBP dB dBBBBP dBP  dB'  
```

auther is program to manage your 2fa (totp and hotp) tokens released under GNU GPL v3 license.

## installation

1. `git clone https://github.com/X3NOOO/auther`
2. `cd auther`
3. `make release`
4. `make install`

or just grab binary from releases

## features

- Support for both TOPT and HOTP
- Encrypted database
- Shell completion (using [cobra](https://github.com/spf13/cobra))

## usage

```ascii
Usage:
  auther [flags]
  auther [command]

Available Commands:
  add         Add token
  completion  Generate the autocompletion script for the specified shell
  get         Generate otp codes
  help        Help about any command
  list        List information
  rem         Remove token

Flags:
  -d, --database string   path to database (default "/home/anon/.auther_db")
  -h, --help              help for auther
      --testing           disable writing to database
  -v, --verbose int       verbosity of output (0-5) (default 3)

Use "auther [command] --help" for more information about a command.
```

## donation

- XMR: `49F3GknYgs7cRfMJghrd9dHZKe63Z6Y3aJKPecDKqLRje5YebzWvz3VWsTa8e8Sk92G7WJEsyp8L1VEeNxmdj2vZNJSACo1`
- DOGE: `DFYc29EsSuSbyLndGrKBGoC2usHRUqiiXb`
- BTC: `bc1q08p6wd86806uf2cj95j4pcgl584jvaqkhs37pp`
- ETH: `0x84FfD8524a66505344A1cbfC3212392Db5b2474d`
- LTC: `Lew3VmzbkaxzoYG3jNHf263oEDMrQ3ecN1`
