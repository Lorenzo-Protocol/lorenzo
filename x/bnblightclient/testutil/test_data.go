package testutil

import (
	"encoding/json"
	"testing"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/types"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

var testHeaders = `[
  {
    "raw_header": "+QM3oEJzCcbdlQcy4HWiWiv5UuGSraw9yhqZLWQhrgRnOt0DoB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlH9fLPGuyDvwx031ZqQap+1l6oTqoP0rid1pu5KJfP2tWeJGjgam0cjRHF+h7/wnBcogggCXoOdfju5EQm7ahwiV0x7d0iM9lFUgC2uDJ4pKg+I+Wu5HoKFY9okKbmIPPpY5rKNbtYt9L+Z5qFJRRw5H7sg2jT4wuQEAAAAEBAAAAAAQAABAAAAAAAEAAAAAAAAQAAAAQAAAAAAAABgCAAoAAEAAAAAAAAAAAAAAAAAAAAAAAAAAECAAAAAAAAgAAAABAAAACAgAAEAgEQAAAAAAgIAAAQAAAAAAAAgAIAIDAAAAABAAAAAIAAgAAAAAIAAAAAAAEAAAAAAAAAgAAAAAAAAAAAAAAAQAAAAEQAAAAABAAAAAAAAAIAAggAAAAAAgCAAAAAIAAAAAACAAABAAgAAAAAAAAAAAIAAIAgAQAIAAAAAAABAAACAAAAAAAAAAABBAAgAAKAAAgAAAAAAAAAEAAAQAAAEBAACQAACAAABIABQAQACAAAKEAoyW84QELB2AgwqRhIRms0RjuQEV2YMBBAyEZ2V0aIlnbzEuMjEuMTKFbGludXgAAGMfg6b4snu4YLV7kp3sJiGu4g8rDYvPUWhV9ZcsOKgex+M5pWiXZ9VaGc/gy6AaR9VkahLeQtkuZRd5BuL3KqPvUIEkpHcT00Sz6Yk4JJgPsfjDzxrmkwQtKWQ+2FFkn/MGaRVzcznPHfhMhAKMlvGgvbPQ4c2wQnwdxTiSLV6bckXsOs/RR3MCAN0yEkYCd32EAoyW8qBCcwnG3ZUHMuB1olor+VLhkq2sPcoamS1kIa4EZzrdA4A6K5D2x4/Uq0SKzvrFiEQA1R3yeqqTWDwZqh7gdp7L7BWoS4sYkOx6752nwWTp2mXqO0UuXxsb7OvxcfigTJuiAKAAAAAAAAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "QnMJxt2VBzLgdaJaK/lS4ZKtrD3KGpktZCGuBGc63QM=",
    "hash": "5pzSSKjPxobzBo8qT8/NTMkSR0DQQla6dVQiyqvOnGo=",
    "number": 42768115,
    "receipt_root": "oVj2iQpuYg8+ljmso1u1i30v5nmoUlFHDkfuyDaNPjA="
  },
  {
    "raw_header": "+QM3oOac0kioz8aG8waPKk/PzUzJEkdA0EJWunVUIsqrzpxqoB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlPmh2w1vIr14/67MvI9HyD35+9vPoFOZS5ckSxTMi/J3by6OEdQqGod3Glt2Ze9IQSnzxzrToLXU2gYIVvDFdRgGM/HstnFTdjgyyt0fMKn5W/oKVQI5oHyfuvfKxyUjuCdYeR/22nFFNfK12/2I6vX9V+GDzxe0uQEAAAAAQAAAAAAAAABAgBAAAAABAAAAAAAQAAABAAAAAAAAgBgCQQIAAAACAAACAAACAAAAAAAAAAAAIAAAAAAAAAAAEAABAAAAAAAACACAAAAgUAAAAIgAgAAAAAAAAAAAAEgAIAoCAAAAABQACIAJCEgAAAIQAAAIAAAAEAAACAAAAAQBAAEAAAgAAAAAAgAAQAEEAAAgAABAAIAAAAEAoQgIgAEAAIAgCiAAAQIAAAAAACAAAAAAAgAAAAAAAAEAAAAIBgAQAIAAAAAAALAAAgAAAAgAAAABkJBEAgAUIAAAggAAAAAAAAEAAAQAABEAAACQACAAAABAQAAAQAgAAAKEAoyW9IQEMEmcgxkwcoRms0RmuQEV2YMBBAuEZ2V0aIlnbzEuMjEuMTGFbGludXgAAGMfg6b4snu4YJBG9SGmRTI9rimtbGBwT9B3RO0vfQR9h1MQEG6\ngShUHYhQyaN0yOuwKAcJMoZ+7GBC/0W300El1XTG60agq+2DFZsVUw8/laVHR1JohprHPJxltabCQKjPanQY6fP1eevhMhAKMlvKgQnMJxt2VBzLgdaJaK/lS4ZKtrD3KGpktZCGuBGc63QOEAoyW86DmnNJIqM/GhvMGjypPz81MyRJHQNBCVrp1VCLKq86caoBj+SMhGe1XVQol2q/3oD3HJYpqNymnCl+E2h/x/HeJFTOPlTtMGezzJCSIkHMkCX9y7nEH/whkpbCgWZ/b9WH2AKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "5pzSSKjPxobzBo8qT8/NTMkSR0DQQla6dVQiyqvOnGo=",
    "hash": "hrYymksY4TjTZRlpSBN0GQta8po7gfsycIKbcASY2d4=",
    "number": 42768116,
    "receipt_root": "fJ+698rHJSO4J1h5H/bacUU18rXb/Yjq9f1X4YPPF7Q="
  },
  {
    "raw_header": "+QM3oIa2MppLGOE402UZaUgTdBkLWvKaO4H7MnCCm3AEmNneoB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlAgmXaAeGmXWK5A8ezTAjLOJvz2ZoM2G9xhAJMTUDEnXnFnhd3KdF+UxmPKkyapJqZmqN98HoFXVT0nGz7vT9LGH2BfuTb7K3bGREf9zILDWvWvdA2+hoLLMqB+ZUb0A/+tThWKYBtDBq9wdnHc/oepZTItAChiZuQEAAAAACgAAAAAAAABAhAAAAAAgAAAAEAAQAAAIAEAAAQAAABAgAAAAAAAgAAAAAAgAAAAACAIAAAAAAEABQCAAABgAAAAAAAAAAAAACwAAAAAhFAAAAoAAAQAAAwAAAAAAAEgBKAJCAAEEAAAAAAAIAAwAQBIQCQAAAAgAEg\nAABAEAAAQAAAABAAAIAgAAAAAAAAAEECAAAAgAARAAAAAAIAIAggAAAAEgCAAAAAIAAAAICiAEAAABAAAAAAAAAUAEACAAAgAAAAAIAAAAAAIAAAgAAAAQAAAQABBMAgAQIAAAEAAAAAAAAAEAAIRAAAEAIACAAEAUAARAAAAAEAAAAAKEAoyW9YQELB2AgweHEoRms0RpuQEV2IMBBAyEZ2V0aIhnbzEuMjIuM4VsaW51eAAAAGMfg6b4snu4YIAbkg33CeYR5ePUr22Y+0B6cXzhzaqm52H406MQpOm2sPqcKFWx9CkCVEjPfMHSgQNybnkvvHkhy1/zmPJ30j1HxMUctQz3whbzJqPD+IIExjXUZ55ITX3MCB8G/R7SO/hMhAKMlvOg5pzSSKjPxobzBo8qT8/NTMkSR0DQQla6dVQiyqvOnGqEAoyW9KCGtjKaSxjhONNlGWlIE3QZC1rymjuB+zJwgptwBJjZ3oDTOBDn6Sbep4/4lvwQQ5Iru61vMy46T3ufx9HIBpuSQjc3qN9yQz967obJ2KAG2MkXLFNs+cZzOJlDOi0vAEPdAKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "hrYymksY4TjTZRlpSBN0GQta8po7gfsycIKbcASY2d4=",
    "hash": "U/mYpWMaVqz/2izhMV/kbJXZt6RH7aq92myHgc0N65A=",
    "number": 42768117,
    "receipt_root": "ssyoH5lRvQD/61OFYpgG0MGr3B2cdz+h6llMi0AKGJk="
  },
  {
    "raw_header": "+QM3oFP5mKVjGlas/9os4TFf5GyV2bekR+2qvdpsh4HNDeuQoB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlBo9nXpxfWTmCIrJN9WqzdPiDKljoFvg9/7KJ\noQwGsXFYD6EZKah0GbiMOL+Uv44iENCpECEoLH+9ZpxtLN3FQ2VFaodGLNInJ4PDRNSqktwfwKulrXDoHkBJeUbfO36W0vP3qXmI6pY9tk4hs7zCfpPxgQPt5abuQEABAQECgAAAAAAAABQhAACAAAgAEAEEAAQAAAIAIAAAUggABgiAAIwAAAgAAAAAAgAAAAACAAAAAAAAEABQCQAgBoAAAAEAAAAAAUAKQCABAAoVAAAAoQIgQAAAgCABAAAgEgAIIpCCAEEABAACIAIBAwgEBKUCQAIAAgIEAAADBgAAAAAAhAhAAgAAgAAAQAAAAEEECAAAAhAAAAAAAAQIAIAkEAAAIEgGgAAAAKAAAAICiAFAAABCAAAAgAAAEAEAKAMBkAQAIIIACAAADIAAgAAAAAQCAAQBBBEAgAQYAAAkAAAAAAAAgEAAARAABEAIACQAGEQAARAAAAAEAAAAAKEAoyW9oQELB2AgxQ/moRms0RsuQEV2YMBBAyEZ2V0aIlnbzEuMjEuMTKFbGludXgAAGMfg6b4snu4YJIPbleNXudqphl3P57Qeze8vIKui9x49Ji+8JvbsfPEQsCEHcOAKiZwU0eM18IAlxmQSMp5ZV6rlamd9YJZtAG1RAr5pVmBH3Z9VYKBB6Linfv7KJ4idJ9fuP2syrcX0PhMhAKMlvSghrYymksY4TjTZRlpSBN0GQta8po7gfsycIKbcASY2d6EAoyW9aBT+ZilYxpWrP/aLOExX+Rsldm3pEftqr3abIeBzQ3rkIAusmYkcDTRNuz3SVgxte5o+H1d3lrHqpXTaad2QSlmLlMEexUawkT71pDpmRMksJFmS3PjOgc40VbfmdgMDzHKAaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "U/mYpWMaVqz/2izhMV/kbJXZt6RH7aq9\n2myHgc0N65A=",
    "hash": "/tffHEYRQNF7S07ZQqJMFPMls1TYuVup7y2wXK1JNXQ=",
    "number": 42768118,
    "receipt_root": "eQEl5Rt87fpbS8/epeYjqlj22TiGzvMJ+k/GBA+3lps="
  },
  {
    "raw_header": "+QM3oP7X3xxGEUDRe0tO2UKiTBTzJbNU2Llbqe8tsFytSTV0oB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlEDTJW6wur6J8OpU7ao5hRMTZhL1oLUrQ7OwaY9QUcEsCey0rI3zQdVEwdgggR3DgEaqXLJPoHDlsf/7d/hR0T7CoQ4BTKFCSjwogNiWxwnXhAVbF62voKC40UPFD3B+CLcNw6CVb0mUyXB7HrDU0yOHE4599D1IuQEABCAADgAAAAAAACBQgQAAAAgoAAAAEAAQAAAAAAAAACAAABQAAAAAAAAAAAAQAAAQCAAgAAAAAAAAAAAAACAAAAAAECAABEACAAAAiAAIACAgXAEAAAAAAAAAAAAAAAAAACgAIAICAAQAAQACAQBYAAwAAAAAAAAABAgAEIAAAABAAAAIAAAASAAAAAAAAEAQAAAMIGAAgAiAAABAAAQAIEIAgAAAAAEiCAIAAAIEAABCCCAAAAAAAgAQIAAAAAIEACAAAgAIAABIAAAAAgQAAAEAAABAICARAJBAAgBgIAAAFAAAAA4AAQEAAAZAAAGACQCAAEAACARACAAAAARABAKEAoyW94QEJ/FkgxrxhIRms0RvuQEV2YMBBAuEZ2V0aIlnbzEuMjEuMTGFbGludXgAAGMfg6b4snu4YJW4rpaKny7/j2J6yF6OGa/8CoJbvlcCJ8aznUGNI/0134dGcUVydLsV4bI/jy+ytAkQnPKry11qU5osPDtSdIWC36Oz+ojwhbFChb7C4uiL7XXQaU3Rn5LClr2LafiekPhMhAKMlvWgU/mYpWMaVqz/2izhMV/kbJXZt6RH7aq92myHgc0N65CE\nAoyW9qD+198cRhFA0XtLTtlCokwU8yWzVNi5W6nvLbBcrUk1dIAQ7fJMzGwX4NwSu+CObGFrXOfBt1Wemsz09sC7Bnqkc2QkXrYKNCS4y9GAKb4dS4+kHMkCUr4DBSax1Lpb7aPXAKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "/tffHEYRQNF7S07ZQqJMFPMls1TYuVup7y2wXK1JNXQ=",
    "hash": "mRmD8fUj+wLdgmjChs9NU4CING9IrQNI47lXajutYSU=",
    "number": 42768119,
    "receipt_root": "oLjRQ8UPcH4Itw3DoJVvSZTJcHsesNTTI4cTjn30PUg="
  },
  {
    "raw_header": "+QM3oJkZg/H1I/sC3YJowobPTVOAiDRvSK0DSOO5V2o7rWEloB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlFM4fzMh/WnR4DC7khIw37GIgmr/oMH3MfErmENdPx0r6h4DsuVe8Rp39J22jLZE2EpctJkWoFqZwltljgLD+fA65TNFJ+9DLCjUfOsMBAAygKnl/TEboBznpcSx7EPB/7uwFoUZG6I3EilSuPEFDlddfQ9Pl0PfuQEABAAAAAAAAAAAAABACAAAAAAAAAAAEAAQAAAASAAAAAAAABASBAAAAACAAAAAAAIQAABAAAAAAAAAAAABACAAAAgAAEAAAAAAAAAACAAAAAAgFAAAAAAAgAAgAAAAAAAAAAgAIAYDAAAAAAAAAAAIAAgAAAgAAAAAAAAAEAAAAAAAAAAAAAAAAIAAAAEAAAAACAAkAAAAAABAAAAAAAAAIAIAgAAAAAAgCAAAAAIAAAAAACAAAAAAAAAAAAAAAAAEAAAIAgAIAIAIAAACAAAAAEAAAAAAAAJAABBAAgAAIAAAkAAAAAAAAAU\nAAAQAAAEAQACAAEAAAABAAAAAAABAIAKEAoyW+IQELBlUgwrfFoRms0RyuQEV2YMBBAyEZ2V0aIlnbzEuMjEuMTKFbGludXgAAGMfg6b4snu4YJC8NDER51yYd9QouMZBWsUkL09XK6HlVS3jBwaWdP9Ltc00m9u+lkWPxHE/ug3IUQHw42ow8Nh8SrC7q0F6ZW0iTIhuvpRMk2/hwZkpibuEZIotiQD9NPBd8xbwTh02J/hMhAKMlvag/tffHEYRQNF7S07ZQqJMFPMls1TYuVup7y2wXK1JNXSEAoyW96CZGYPx9SP7At2CaMKGz01TgIg0b0itA0jjuVdqO61hJYBgOGgIb8Li1e/V9JK+VkxmSfxOV3IPdHnuJO2wjZnM5iS2ZDDv4o6DNveBpOyZWqI18YhsC3btDSI+HMatW6ERAaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "mRmD8fUj+wLdgmjChs9NU4CING9IrQNI47lXajutYSU=",
    "hash": "oVNGVaqjKSwk3SFXDPKSCsQUFOT5x4TvLjkVGNcv0cY=",
    "number": 42768120,
    "receipt_root": "HOelxLHsQ8H/u7AWhRkbojcSKVK48QUOV119D0+XQ98="
  },
  {
    "raw_header": "+QM6oKFTRlWqoyksJN0hVwzykgrEFBTk+ceE7y45FRjXL9HGoB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlHbXbuiCPeUqGkMYhMLKkwxecr/zoFE/ItCJbC8xXn8METNhmPPYzaM/jkGXlHqcJS7/HJy4oCSZ42wFvoCR1YKK10r1p7I1USeDn7NshosX/KQvVJzCoEdaXuZq23jUBxWoKQhaEooXUjXhr6GaVxr5BSE/jA9YuQEABIAAAAAAAAAAAABEAA\nAAAAgAAAAAEAAUAAAAAAAAAAAAABACQAAAACAAAAAAAAAACAAAAIAAAAAAAAAAACAAAAAAEAABJAAAAAAACACIAAAgFAAAAAAAgCKAgAEAAUAAAAgAIAoCAAEAAAAAAIAIAAoQBAAAAAAIBAAAEAAQAAAAAQBAAEAAAAAAAQAAAAAAAAEEAAAAAAJAAIABAAAAMAIEgAEAAIAgCAAAAAIAAACCACAAAAAAAAAAAAAIAAAEAABIBgAIAIAIAAAAACAAAgAAAIAAAAAAABBAAiAgIAAAkAAAAAgAAAEAAAQAAAEAIACAgEAAAABAAAAAQABAAAKEAoyW+YQELB2AgxNqNYRms0R1uQEV2YMBBAyEZ2V0aIlnbzEuMjEuMTKFbGludXgAAGMfg6b4snu4YJPzjDT0GRjIpqUJYy7joRPEy+6oY/pe88N7O4UEUJ0M8OQNQqcekpUbFQam50aJfBOzjZhbtny7a4zi1N84TtmVCD9U/Gg+F1Zzsc3s9/sq/9LGHq+QQAr/IKsaHhxHvPhMhAKMlvegmRmD8fUj+wLdgmjChs9NU4CING9IrQNI47lXajutYSWEAoyW+KChU0ZVqqMpLCTdIVcM8pIKxBQU5PnHhO8uORUY1y/RxoDTGvibd+6gCqgaZehmHkxRVl+0Xcv32P3xtyZIJTF/NnfD3hHIEyPWO7wiSKgzZ3T2GHEOg0CuIXu/1gjyQ1cBAKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGDAgAAgA==",
    "parent_hash": "oVNGVaqjKSwk3SFXDPKSCsQUFOT5x4TvLjkVGNcv0cY=",
    "hash": "vev78tC4b4tzzDHiEeR8w4nyEi4zjFBIQI74AVIpwyE=",
    "number": 42768121,
    "receipt_root": "R1pe5mrbeNQHFagpCFoSihdSNeGvoZpXGvkFIT+\nMD1g="
  },
  {
    "raw_header": "+QM6oL3r+/LQuG+Lc8wx4hHkfMOJ8hIuM4xQSECO+AFSKcMhoB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlH9fLPGuyDvwx031ZqQap+1l6oTqoDjYMNlXBk1gfSezWJHiDpCdYZFZOzhgDSOTFWZF4Hk5oBGSzfXouXvsPk0vqBxod4BDKDXj26e3vOYL4IQVWreeoKq7vKX5aTwFWKS9M1NMb9AvNm1SMvpVK+L15WGKYwa3uQEABAAEAAAAAAAAAABAAAAAAAAAAAAAEAAQAAAAAAAAAEAAABASBAAAAACAAAAAAAAIAABAAAAAAAAAAAAAAAAAAAAAAEAAAAABAAAACACAAAAgFAAAAAAAgAggAAEAAAAAgAgBIAoCAAAAgAAAAIAIAAgAAIAAAAAIAAAAEAAQAAAAAAgAAAAAAAAAAAAAAAAAAAEEAAAAAABAABAAAAAAoAAAggAAAIAgCAAAAAIAAAAAACAAAAAAAAAAAAAIAAAEAAAIDgAIAIAIAAAAACAAAgAAAAAAAAAAABBAAgAAIAAEkAAAAAAAAAUAAAQAAAEAQACAAECAAABAAAAACABAIAKEAoyW+oQELB2AgxHFX4Rms0R4uQEV2YMBBAyEZ2V0aIlnbzEuMjEuMTKFbGludXgAAGMfg6b4snu4YIyfc/s+p/et/iWU+THzOvy9B0ySp6USXdpxIKxbTK3l9+/uoIe4awDuQ4Es1RxUUgZvuqgoNlojfiGCkX7juFrxId09xacjEDF0gZYks4cuY7c5+N15NHMMrRloC2rOwvhMhAKMlvigoVNGVaqjKSwk3SFXDPKSCsQUFOT5x4TvLjkVGNcv0caEAoyW+aC96/vy0Lhvi3PMMeIR5HzDifISLjOMUEhAjvgBUinDIYDhrmh+IwqSmWWgvSREkZJ8xqVYS2dY692iGKrMHCyjWkvbpzmmv8lbCTva14R8gPmv36ZmO+99TDnwLPDBM20YAKAAA\nAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGDAgAAgA==",
    "parent_hash": "vev78tC4b4tzzDHiEeR8w4nyEi4zjFBIQI74AVIpwyE=",
    "hash": "DAsmidyvJU+ODhrSjT4skBCp7ZMHf+1SAgdMEK9VP18=",
    "number": 42768122,
    "receipt_root": "qru8pflpPAVYpL0zU0xv0C82bVIy+lUr4vXlYYpjBrc="
  },
  {
    "raw_header": "+QM3oAwLJoncryVPjg4a0o0+LJAQqe2TB3/tUgIHTBCvVT9foB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlPmh2w1vIr14/67MvI9HyD35+9vPoH+RfV/bNu3peJU2dBiY1GvW5E2cnymNa5iW9GsKTMMRoM2Juo7a3UN4O6OfhkI4AA0QaCwCStDKloi3bBMWF8X8oLRur5PilCidO+cDEHJyMfEdNobXclN7h5IcMqHUjshmuQEAAAAFTAAAAAAQAABABAAAAAgAAEAAABAQAAAIAAAAAAAgABAiAAAAAAAiAAEAAAgACAAAKAAAAAAAAEABACAAABgAgAAABAAAAAAACQAkCAAgEEAAAgAAgYAAAgEEAAAAABgAIAJCAAEEAAAAAAAIAAgAEBAAAQIAAAAAEAAQBAAABAAAAACBAAAAAgAAAQQAIAAEFCAAAABAAAAAAAAAIAIggAAAAAFgCAAAAAIAAAAICiAEAAABgAAAAAAMAEAAIKAqAgAAAIYIAAAAAAAAACAACAAQAAAAABBAAgAAIAAAlIAAAAAAAAEAgAQIAAEAIACEAEAQAARIQAAAUAgAAAKEAoyW+4QEMEmcgxFStYRms0R7uQEV2YMBBAuEZ2V0aIlnbzEuMjEuMTGFbGludXgAAGMfg6b4snu4YIFVpwUDkSkGKSUF4VdwL1rtu8jV\nqKvpj5+z2Yrx4XTkdDV+72ILinFbKW3+edDtXAV7m1pTN7PldpviQmmFUGwZhTW8rRCcMtOpNbwMxmqvL0hiDnHlKpCXYpqiCiJvh/hMhAKMlvmgvev78tC4b4tzzDHiEeR8w4nyEi4zjFBIQI74AVIpwyGEAoyW+qAMCyaJ3K8lT44OGtKNPiyQEKntkwd/7VICB0wQr1U/X4D6+7aTNc7dC80u8yCEcrO0tfEvvMuXVwOq0+HRRfnL+2V/iuGrdiGIFo6Uhr9LsBjNip2U59y2MCBqq1N540W+AaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "DAsmidyvJU+ODhrSjT4skBCp7ZMHf+1SAgdMEK9VP18=",
    "hash": "7YIw2+1/P+TxeKWlo8Hz6eXn2qlCokjWXHB18PjIbHQ=",
    "number": 42768123,
    "receipt_root": "tG6vk+KUKJ075wMQcnIx8R02htdyU3uHkhwyodSOyGY="
  },
  {
    "raw_header": "+QM3oO2CMNvtfz/k8XilpaPB8+nl59qpQqJI1lxwdfD4yGx0oB3MTejex116q4W1Z7bM1BrTEkUblIp0E/ChQv1A1JNHlAgmXaAeGmXWK5A8ezTAjLOJvz2ZoApucgdYivySXOBrXvLYR472JFJucpL3rawTYKYCqG/KoFP83ugVZm/BgKzV0JM2/8dJ9h5Jm3PY3+fLmkSGO8vgoA/il3xOBUkS6hvhS0TsPyRwvMtondtX5EE9f9iSDqQNuQEABQAABAAAAAAAAABADAAAAAAAAAACAAAgAAAICAAAAAAAABggAAIAAAAgAAAAAAgQACAACAAAAAAAAEIBACAAABgAgAAAAAAAAAAACQAAAAAhEAAAAgAAAQAAAgAEAAAAAAgAIABCAAEEABAQAAAAAKgAABg\nAAQAQAAAAEAAABAAAAAAAAQABAAAAAgEAAAAAKAAEECAAAAAAAAAAAgAAIAIAAAAAAAAgCAABAAIAAAAIAgAEAAABAAAAAAAEAEAAAAAAAgAQAAAAAAACABAAAAAAAAAQAAAAABBIAgAAAAAAEIAAAAAAAAEIAIQAAAEAIACQAACQAARAAAAAEAAAIAKEAoyW/IQELB2AgwcNpIRms0R+uQEV2IMBBAyEZ2V0aIhnbzEuMjIuM4VsaW51eAAAAGMfg6b4snu4YJZNzT/r8wavZcHSuEtn/0EM37n9wPXusnomyjo6LMBpTO8Ec9/zcXPl10fUnyF8fgedOfZR6zkz9XCUDXhKLE7xJh+931/6seExjAMI7cyBWpn4P2iPn0xlX2UnZu3OdfhMhAKMlvqgDAsmidyvJU+ODhrSjT4skBCp7ZMHf+1SAgdMEK9VP1+EAoyW+6DtgjDb7X8/5PF4paWjwfPp5efaqUKiSNZccHXw+MhsdIAx9obtXZrV61lCwpwhxwrrULcxGLG1KhxNqlervzBfxDVyxfB1sPBE1J4p2EKerRClMQrUS1yzwmshAAQZrx9tAaAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIgAAAAAAAAAAICgVugfFxvMVab/g0XmksD4bltI4BuZbK3AAWIvteNjtCGAgA==",
    "parent_hash": "7YIw2+1/P+TxeKWlo8Hz6eXn2qlCokjWXHB18PjIbHQ=",
    "hash": "AapiOfFbavBdm8vyJsKTSoJ6pC7LItyN/Mfy2lIKLU8=",
    "number": 42768124,
    "receipt_root": "D+KXfE4FSRLqG+FLROw/JHC8y2id21fkQT1/2JIOpA0="
  }
]`

var receiptRLP = `b9034802f903440183113d2db9010000000000000000000000001000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040080000000000000000000000008000000000000000000840800000000008004000000400000020008000000000000000800002000029000000000000010000000000000000002000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000004000400001060000000000000000000000000000000000000000000000000000000000000000000f90239f89b9449ff00552ca23899ba9f814bcf7ed55bc5cdd9cef863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01a00000000000000000000000004730ddefe385300417cb5617f3adf0105aff6806a00000000000000000000000000000000000000000000000001bc16d674ec80000f89b942a45de58552f2c5e0597d1fbb8ec83f7e2ddba0df863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01a00000000000000000000000000000000000000000000000001bc16d674ec80000f8fd949adb675bc89d9ec5d829709e85562b7c99658d59f884a065199e8f9cbc273da46a3e57ea48d1ba2ced8379e8f19785847bb3bf1f712dfea00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01b86000000000000000000000000049ff00552ca23899ba9f814bcf7ed55bc5cdd9ce0000000000000000000000000000000000000000000000001bc16d674ec800000000000000000000000000000000000000000000000000001bc16d674ec80000`

var proofRLP = `f909190ab9034802f903440183113d2db9010000000000000000000000001000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040080000000000000000000000008000000000000000000840800000000008004000000400000020008000000000000000800002000029000000000000010000000000000000002000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000004000400001060000000000000000000000000000000000000000000000000000000000000000000f90239f89b9449ff00552ca23899ba9f814bcf7ed55bc5cdd9cef863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01a00000000000000000000000004730ddefe385300417cb5617f3adf0105aff6806a00000000000000000000000000000000000000000000000001bc16d674ec80000f89b942a45de58552f2c5e0597d1fbb8ec83f7e2ddba0df863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01a00000000000000000000000000000000000000000000000001bc16d674ec80000f8fd949adb675bc89d9ec5d829709e85562b7c99658d59f884a065199e8f9cbc273da46a3e57ea48d1ba2ced8379e8f19785847bb3bf1f712dfea00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01b86000000000000000000000000049ff00552ca23899ba9f814bcf7ed55bc5cdd9ce0000000000000000000000000000000000000000000000001bc16d674ec800000000000000000000000000000000000000000000000000001bc16d674ec80000f905caf887ac6551456c355274383766706253382f657065596a716c6a32325469477a764d4a2b6b2f4742412b336c70733dac6b303866774153725246654c746e6f6f3347436e737346466170724f447753645363674f616a4c6f4b62413dac4e347862563344515a725a2f4c4a335138735368796175635032564a67702f457a73756d722f4a6e776c773df9053eb853f851a0934f1fc004ab44578bb67a28dc60a7b2c1456a9ace0f049d49c80e6a32e829b080808080808080a0e44c574445030e690f39db1bde66164c3ed162119f3d3a89fbf80b0a1e55e53e8080808080808080b90194f9019180a01b5a6c3ba1663a9d926c26525bd5af6f01a683ce86cc5eea866c74d48d12b5b7a06972bc747931f7c2861a86b54b86cc82b312a78f34b7d129238e2a41bad0f600a0154cd2dc1505aa39d816623f1a1d0c89f75dece7893404cc5e0b2d1aa1ee19aea0efccd91cbb07756c8e754b3a46fa9fd9dbead952bf34e413a61a1238f59d2e22a0c5ebc118bf68a84b269acf053efb6f986e03df62e64301a030095000c0542f52a0fa6af3a5cd5a19b4f5d24461d9ea4052c23e5c53536f9989e925224a35d48231a0b6429abd7c1b16888eeae2ae2c53ec63bc6a5cfd0606da5f5b0ddf2195fc97e0a0e3d3e0f89776bb9aa145ffa0f1e5dc39d777d2c9745ea5b121b8c0e979155732a085e86d11a1376c7e25ed78d95eb63a8ee8b82639c16c6aee79292a884658c042a0378c5b5770d066b67f2c9dd0f2c4a1c9ab9c3f6549829fc4cecba6aff267c25ca06416535922d9bcd406243511108208059f1535735b5ee8fff3a99f0939fbd8c0a071085c605c2e76d176d7eb8e99b72d8af248f7e88e53ac1ca232798709a99ca680808080b9034ff9034c20b9034802f903440183113d2db9010000000000000000000000001000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040080000000000000000000000008000000000000000000840800000000008004000000400000020008000000000000000800002000029000000000000010000000000000000002000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000004000400001060000000000000000000000000000000000000000000000000000000000000000000f90239f89b9449ff00552ca23899ba9f814bcf7ed55bc5cdd9cef863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01a00000000000000000000000004730ddefe385300417cb5617f3adf0105aff6806a00000000000000000000000000000000000000000000000001bc16d674ec80000f89b942a45de58552f2c5e0597d1fbb8ec83f7e2ddba0df863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01a00000000000000000000000000000000000000000000000001bc16d674ec80000f8fd949adb675bc89d9ec5d829709e85562b7c99658d59f884a065199e8f9cbc273da46a3e57ea48d1ba2ced8379e8f19785847bb3bf1f712dfea00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000dfb41dc2173d2be024e6d64a83fd011d4ae43e01b86000000000000000000000000049ff00552ca23899ba9f814bcf7ed55bc5cdd9ce0000000000000000000000000000000000000000000000001bc16d674ec800000000000000000000000000000000000000000000000000001bc16d674ec80000`

// ProvedReceipt is the struct for proving receipt
type ProvedReceipt struct {
	Receipt *evmtypes.Receipt
	Number  uint64
	Proof   *types.Proof
}

// GetTestHeaders returns the test headers
func GetTestHeaders(t *testing.T) []*types.Header {
	var headers []*types.Header
	err := json.Unmarshal([]byte(testHeaders), &headers)
	require.NoError(t, err)

	return headers
}

// GetTestProvedReceipts returns the test proved receipts
func GetTestProvedReceipts(t *testing.T) *ProvedReceipt {
	receipt, err := types.UnmarshalReceipt(common.Hex2Bytes(receiptRLP))
	require.NoError(t, err)

	proof, err := types.UnmarshalProof(common.Hex2Bytes(proofRLP))
	require.NoError(t, err)

	return &ProvedReceipt{
		Receipt: receipt,
		Number:  42768118,
		Proof:   proof,
	}
}
