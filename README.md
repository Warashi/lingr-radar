lingr-radar
===========
コンフィグファイルは以下の順に探します

1. $XDG\_CONFIG\_HOME/config.json
2. $HOME/.config/lingr-radar/config.json
3. DIR in $XDG\_CONFIG\_DIRS, $DIR/lingr-radar/config.json
4. $HOME/.lingr-radar/config.json

コンフィグファイルのExample
```json
{
    "User":"your username",
    "Password":"your password",
    "APIKey":"your apikey"
}
```

----
source の一部は lestrratさんの pecoからコピーさせていただきました.
