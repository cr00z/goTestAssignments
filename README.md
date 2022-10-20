# Задание

Написать скрипт, который успешно проходит опрос, расположенный по адресу: http://185.204.3.165

Скрипт должен быть написан на языке программирования GO.

В настройках запуска скрипта должна быть какая-нибудь возможность указывать количество параллельных потоков для независимого выполнения сценария (эмуляция нагрузки).

Не нужно тестировать указанный сервер под большой нагрузкой! Максимум 10 RPS. Мы будем анализировать только код вашего скрипта.

Скрипт должен быть помещён на любой публичный GIT репозиторий. Тестовое задание присылается на проверку виде URL ссылки на GIT репозиторий.

Где-то в корне репозитория должна быть инструкция по запуску скрипта (в любом формате и форме).

# Решение

Конфигурация в секции const

- workNum - количество задач
- RPS

Запуск: `make run` или `go run main.go`

Пример запуска при workNum=10, RPS=5

```
➜  goTestAssignments git:(main) ✗ go run test.go
2022/10/20 14:32:31 Work 1 step 0 started
2022/10/20 14:32:31 Work 2 step 0 started
2022/10/20 14:32:31 Work 1 auth complete: [SID=fAdEVjc3QoUzJbduG3MYufR5kEYJggzE; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:31 Work 2 auth complete: [SID=BYvf7iDsKxFmdHYdEvxULbwrqD34z9xr; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:31 Work 3 step 0 started
2022/10/20 14:32:31 Work 3 auth complete: [SID=ewxc1jg7Qn0hCbq5TxC98D4JRvjGXsYW; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:31 Work 4 step 0 started
2022/10/20 14:32:31 Work 4 auth complete: [SID=cdmLY9mC5MkxgrRNZ2IEcxfo2SdESh7a; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:31 Work 5 step 0 started
2022/10/20 14:32:32 Work 5 auth complete: [SID=lmosVz7fL3CB0QFCNimfizziibf3IiKr; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:32 Work 6 step 0 started
2022/10/20 14:32:32 Work 6 auth complete: [SID=tFIhdvUwP401TxgHZfTTm9KGwxxgYyAb; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:32 Work 7 step 0 started
2022/10/20 14:32:32 Work 7 auth complete: [SID=22GS0yjHuxGPdruIheuaAsrtrpYl8T9R; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:32 Work 8 step 0 started
2022/10/20 14:32:32 Work 8 auth complete: [SID=DntDAPAINkor7Z8GW5kyr5LgcXF83njW; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:32 Work 9 step 0 started
2022/10/20 14:32:32 Work 9 auth complete: [SID=065AdDb3NSpCelYkTJiblBTja6F1xrLl; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:32 Work 10 step 0 started
2022/10/20 14:32:33 Work 10 auth complete: [SID=tBtZKcM4EKvzS2egt0dVOkaYUNtMfWPd; Max-Age=3600; HttpOnly; SameSite=Strict]
2022/10/20 14:32:33 Work 1 step 1 started
2022/10/20 14:32:33 Work 1 question 1 answer: map[xvBKpERtUaQg1MZv:JWQBxWDaZRY3JPXQ]
2022/10/20 14:32:33 Work 2 step 1 started
2022/10/20 14:32:33 Work 2 question 1 answer: map[12dnc5Tx3SyC3FCd:DDNeSrv0GPZN8hw iP5DTOoUjaSwoLii:EaedR3gowGdrL]
2022/10/20 14:32:33 Work 3 step 1 started
2022/10/20 14:32:33 Work 3 question 1 answer: map[N0pBn9dPQniVMA6k:YC5YlEV4jgyydi S9xvYsjanWFErq1x:UdYYy9GLZG8Ao i5bbq6gHeQGQjcke:test qwL37mr92MJKCsVr:bkraGI9rKvJ3Z]
2022/10/20 14:32:33 Work 4 step 1 started
2022/10/20 14:32:33 Work 4 question 1 answer: map[8jslcCpaaeUW7O7P:test 9S03pWwX49pyIOsJ:L4bBSxxCA8A4O Rrr0SxmWIHTdFZqd:test]
2022/10/20 14:32:33 Work 5 step 1 started
2022/10/20 14:32:34 Work 5 question 1 answer: map[O3lPjAwtBcKCUIWI:MeyquLrQWNLwmuKm]
2022/10/20 14:32:34 Work 6 step 1 started
2022/10/20 14:32:34 Work 6 question 1 answer: map[AnNhEasFT7oB64b1:U0211e8wOQhmAoh Jrf8bw3JNJUrOapN:ngDuXP1wRFmXa KhnC6dSXGPrdfdcN:b6X6aLWs7lkWn4gT]
2022/10/20 14:32:34 Work 7 step 1 started
2022/10/20 14:32:34 Work 7 question 1 answer: map[qTJive4kr5ktu5zU:w1AbkB6aHzW9W tTy8QIbzDhq8pVzK:WbzKdn77YjJEp xGcNHZDV9p9ebRlq:RaqxeyLpVWdFk]
2022/10/20 14:32:34 Work 8 step 1 started
2022/10/20 14:32:34 Work 8 question 1 answer: map[3rOdW7D1ou9nQjrd:test]
2022/10/20 14:32:34 Work 9 step 1 started
2022/10/20 14:32:34 Work 9 question 1 answer: map[8uXg6Dt8RnnYCOY2:test C2EWPZSWfRn9G9F8:NwrJLysqCZHQo cjgdtMNCPbQ61yhI:8qhAusbi2PSVph9 vCF5rNDkwffdUtr4:test]
2022/10/20 14:32:34 Work 10 step 1 started
2022/10/20 14:32:35 Work 10 question 1 answer: map[M7JFwVeSTqf3uA6N:UjjVzAQBoyrGL amtfKsIRyPMq1dBo:V1RY683qvwGQfBa jzAKv9xYfx32dr64:PmrrAGvmPXAU6]
2022/10/20 14:32:35 Work 1 step 2 started
2022/10/20 14:32:35 Work 1 question 2 answer: map[ti1I1OwybwG3RLXu:ZSPhTPjGgzLxVZL]
2022/10/20 14:32:35 Work 2 step 2 started
2022/10/20 14:32:35 Work 2 question 2 answer: map[MfHZeIiEUxNPDZbW:WDsf13nqwvqLl44D XQOrB4eHBaL0NtvJ:3ccnU1TMNVtUg fIMV4hFlZL8Okty1:28slitsWwaeFlfP7]
2022/10/20 14:32:35 Work 3 step 2 started
2022/10/20 14:32:35 Work 3 question 2 answer: map[QyZ38LueNSkXNoVM:test WZ5nyx5QatbuxQmx:Ibdb7YnLCafS5cDz]
2022/10/20 14:32:35 Work 4 step 2 started
2022/10/20 14:32:35 Work 4 question 2 answer: map[LYOc5hUh3R6vabQg:x3bK5JeKhMEABY rDmAUyalqvRwWGJ8:my9VLCdkbhdsj]
2022/10/20 14:32:35 Work 5 step 2 started
2022/10/20 14:32:36 Work 5 question 2 answer: map[3lBxMbvlSMvEPWI5:LMLVi1YWch38B 5gL6o7ZzSoGkiygd:0xNOfqqfiH7Rof tkWPtuHH4H9onrL3:bPK5pYgyuUuls yXVtb17TOeRMhgYS:test]
2022/10/20 14:32:36 Work 6 step 2 started
2022/10/20 14:32:36 Work 6 question 2 answer: map[RobYOI3SYy2IXSd6:3638JkqdBilncV lXmn7UiYcYxyqO5D:uLMMYXrI6oTes0]
2022/10/20 14:32:36 Work 7 step 2 started
2022/10/20 14:32:36 Work 7 question 2 answer: map[2aCh9hAwkWfAUJac:ZGncdpigjzMPftCr 465JX7yxL5SM1nAj:test Gq0zx5T6miYmyfei:test]
2022/10/20 14:32:36 Work 8 step 2 started
2022/10/20 14:32:36 Work 8 question 2 answer: map[2Tef2pnHiEG1VAo0:uEeKVTacPsktN5C y7lJTKWNTCIiTHli:1agUI12vh8XI84hN ySA8XQEEsMhuqvN5:GXelrnOg6CBus]
2022/10/20 14:32:36 Work 9 step 2 started
2022/10/20 14:32:36 Work 9 question 2 answer: map[IGsrCa4ptAXqLJZr:y95TflpnAZv18 a2dh4wdj7Q695ZhA:J0CsO1qZKJM60 ebDMxvxjL8tVoTgW:ZJc1L2FvbfnduhN]
2022/10/20 14:32:36 Work 10 step 2 started
2022/10/20 14:32:37 Work 10 question 2 answer: map[Nkj66fJjZQuvJhdx:test]
2022/10/20 14:32:37 Work 1 step 3 started
2022/10/20 14:32:37 Work 1 question 3 answer: map[55wLeDcG5zGd0dRy:QEupXO1JgL2cUWz fcjYDzl1EqJNVZU8:TJLEMZEWnCRfg pGvKSqFa4pzBXC4t:XfzJvC59mNFT1U]
2022/10/20 14:32:37 Work 2 step 3 started
2022/10/20 14:32:37 Work 2 question 3 answer: map[33aURvXSpVVKPbRC:wlbnb5WDX6lNv UCf07mQUZZ6N6jGC:Zu5XzmvPPNFpP0i]
2022/10/20 14:32:37 Work 3 step 3 started
2022/10/20 14:32:37 Work 3 question 3 answer: map[GuacFMmkO3aFYcSh:IWVtnJn35RKNQ gq64gKGAY5iDIyQw:test]
2022/10/20 14:32:37 Work 4 step 3 started
2022/10/20 14:32:37 Work 4 question 3 answer: map[SN2Yv5ycqIeQBj7t:test XA3Qapxy52NFiT5m:5y49c65QieMCIn wO7RylEI8Yde2iXF:Jk07ORxlg1Ylwqtw waevoKWoxuo5kjpK:test]
2022/10/20 14:32:37 Work 5 step 3 started
2022/10/20 14:32:38 Work 5 question 3 answer: map[4xL3aipk57OT9NPN:test YxMcbmy026Gveow1:test nbEnOlrqpKoQhrHw:GzKQYTeQubKYIzN]
2022/10/20 14:32:38 Work 6 step 3 started
2022/10/20 14:32:38 Work 6 question 3 answer: map[KMKPB0o1RhGLPi1T:blFxkmWvC88a5 MEhn4R1n8FLvikI7:8ndTXBg5v3b4y tlNx0T37KihT3ihm:Gt51OzhsOteiB9]
2022/10/20 14:32:38 Work 7 step 3 started
2022/10/20 14:32:38 Work 7 question 3 answer: map[nUDP7nN0ajArVeOI:XbBJN9K8IeCJO]
2022/10/20 14:32:38 Work 8 step 3 started
2022/10/20 14:32:38 Work 8 question 3 answer: map[au0pxtOTp6fNu3RY:e5YZa0XJPUYlH4O]
2022/10/20 14:32:38 Work 9 step 3 started
2022/10/20 14:32:38 Work 9 question 3 answer: map[5lI9LIIRsfeHOSnZ:motI3Nu3Zl2rIF]
2022/10/20 14:32:38 Work 10 step 3 started
2022/10/20 14:32:39 Work 10 question 3 answer: map[1QEw0zmrs8fCvCfG:E6cdWUuOsbdu98Y ELInQ7yBbCD1hPKU:M6KRTXb4WuNX5Lg Z3YfTqMv1rvHilI2:n4RBxG5Z9kYut l7WgHzNmSWXleMFf:test]
2022/10/20 14:32:39 Work 1 step 4 started
2022/10/20 14:32:39 Work 1 question 4 answer: map[11EQrFEEdUjbM1n1:aH5RB6RxVtnIrq g4ESgWbkxphs0hV3:kbUh667Fh7SFG]
2022/10/20 14:32:39 Work 2 step 4 started
2022/10/20 14:32:39 Work 2 question 4 answer: map[4c7DQNSP8xLLfJqF:js25kmyBG4BSit YhkyRva0BYkysJUO:AyYQKjX2AMSCRsiH w2DeqgCuaHE525ot:mRRtjHbm8NaWqF]
2022/10/20 14:32:39 Work 3 step 4 started
2022/10/20 14:32:39 Work 3 question 4 answer: map[6JSkENo12ie6FlO0:yauw631QhV4wvRDf]
2022/10/20 14:32:39 Work 4 step 4 started
2022/10/20 14:32:39 Work 4 question 4 answer: map[2lDqdLN9uROVWohL:XN30QxRaOwh6RlM 44nkhZLkybh3y0xP:nsGN8huDBO8wp8N nM6n2ksvQHlXMlHB:test wdhwCjvkG4TOTdyR:BBF0N07hz7sbGdf]
2022/10/20 14:32:39 Work 5 step 4 started
2022/10/20 14:32:40 Work 5 question 4 answer: map[38fSIrKSCZIKPPXR:osmxlQSvy6utlS Exj5kWm1RDyviD0z:rGmG85pVXZjbTL1 a8XsUbdPZxabGaeS:bMVo2PVInlPO3]
2022/10/20 14:32:40 Work 6 step 4 started
2022/10/20 14:32:40 Work 6 question 4 answer: map[3RAwjBVItq1wu4a9:test ZcHOtUiRGeB6kqXY:j6mbYRXFrw9LE d6NVZcGHtrWNVdxp:04HTU5iyN9s210gJ]
2022/10/20 14:32:40 Work 7 step 4 started
2022/10/20 14:32:40 Work 7 question 4 answer: map[31MmJIrlstxx9DNM:fIbpkmBBLq09g 7C8zY0TLbLZoY2le:test UAwPh9usanitehoN:test]
2022/10/20 14:32:40 Work 8 step 4 started
2022/10/20 14:32:40 Work 8 question 4 answer: map[hQVYDRoY4rUE8ztR:bz51ZwAWvyTEyVwC]
2022/10/20 14:32:40 Work 9 step 4 started
2022/10/20 14:32:40 Work 9 question 4 answer: map[FEAiplpTsmGELnXq:SjW906rC4hxAl eTJQNa4JQD34lmd8:N8cuSt1WkxcjNjM8]
2022/10/20 14:32:40 Work 10 step 4 started
2022/10/20 14:32:41 Work 10 question 4 answer: map[FsWhdVLhoYnUKwCb:rgQlfkuUJZU5fvk]
2022/10/20 14:32:41 Work 1 step 5 started
2022/10/20 14:32:41 Work 1 question 5 answer: map[WL2u4SxbzfVVIRC3:FHsqAPyPUrCc2Z bSTUBNXVjdyqyr4y:8xZhuy7EEc6M8Iz4 qTuSRVKjKzQTGh31:DNN3KhPFzg5h7 vCHBRUb1TKDA9UTX:7el7eZiv11PPBqr]
2022/10/20 14:32:41 Work 2 step 5 started
2022/10/20 14:32:41 Work 2 question 5 answer: map[KS8MWURp9kUhYh3e:test]
2022/10/20 14:32:41 Work 3 step 5 started
2022/10/20 14:32:41 Work 3 question 5 answer: map[M1Cs5flLElXdxdio:test i24jkhicnqv0Fadj:test yTfXovMKOKGUWU4b:kKI1AMw7z0vX2eU]
2022/10/20 14:32:41 Work 4 step 5 started
2022/10/20 14:32:41 Work 4 question 5 answer: map[7HwO0SesD3QwkcQQ:test HfCttJEzzdRZfEpr:TKXj42w4IdvEr hpuh9ktI0K8dxHnB:LsnaKqIWxEWb40Vd oApOfT49fMgArwbu:CnJqIGqFqyR61KQc]
2022/10/20 14:32:41 Work 5 step 5 started
2022/10/20 14:32:42 Work 5 question 5 answer: map[69AEkl4miGmqU1ip:test QUeWLjRrExYVMQpC:yHSbr9z1K4HbNS]
2022/10/20 14:32:42 Work 6 step 5 started
2022/10/20 14:32:42 Work 6 question 5 answer: map[G2Eo4Jl47R1NnRAx:test MfvsneOTtLzFIlcB:zoyoKn1LJgroT]
2022/10/20 14:32:42 Work 7 step 5 started
2022/10/20 14:32:42 Work 7 question 5 answer: map[loqdAjQ4sHAUD653:Yvpo7VY5HIpszC7]
2022/10/20 14:32:42 Work 8 step 5 started
2022/10/20 14:32:42 Work 8 question 5 answer: map[JyZJ3j0wcwSUI9WT:test VKeCNDf2nt7m06YF:2sAXcqhnNWAMIJm hoPwSAabxV1DwoE8:test]
2022/10/20 14:32:42 Work 9 step 5 started
2022/10/20 14:32:42 Work 9 question 5 answer: map[BrCff3Z1JBom0oL3:test FNN3qxM9uOUs6Wd6:vM3VggpbQmNkw]
2022/10/20 14:32:42 Work 10 step 5 started
2022/10/20 14:32:43 Work 10 question 5 answer: map[EUHk7IBgrUHWWcBd:T763d4qPYrm6HMvU XghjM0Lbm4u2RNbX:JdLzWB7qvb7Pn4 jvGq0KE3DMxiYLBV:test yL2h3UNl3p87Emur:test]
2022/10/20 14:32:43 Work 1 step 6 started
2022/10/20 14:32:43 Work 1 question 6 answer: map[CBmiZzh7bPd6nTBY:CsookKMrId7b2d w6V8izpietOLIr5p:test]
2022/10/20 14:32:43 Work 2 step 6 started
2022/10/20 14:32:43 Work 2 question 6 answer: map[TtLRYhCvYa4tnZ9A:test]
2022/10/20 14:32:43 Work 3 step 6 started
2022/10/20 14:32:43 Work 3 question 6 answer: map[FOFZaqgSnfa1dZjz:test NKaytKkjZ1hTMmZE:6TUuD4pZrKmrhMY]
2022/10/20 14:32:43 Work 4 step 6 started
2022/10/20 14:32:43 Work 4 question 6 answer: map[icVw8EG1pcwLhU37:cpnOtGHeZ6e1UD zi1Z71H4lcObbPY3:SZrOtni1DUqWCJg]
2022/10/20 14:32:43 Work 5 step 6 started
2022/10/20 14:32:44 Work 5 question 6 answer: map[3NYI705rGUopqBmZ:test hDLSeskW1uJe33Db:test iSfHVr8MxsApW3jk:duFQhm5ayDKyZuo]
2022/10/20 14:32:44 Work 6 step 6 started
2022/10/20 14:32:44 Work 6 question 6 answer: map[aIwe8AfOMBxwrVLd:test]
2022/10/20 14:32:44 Work 7 step 6 started
2022/10/20 14:32:44 Work 7 question 6 answer: map[9FXeD1HaoHb2ygP7:RTvT7HLvRK77gAi PzF490NwretuisUV:test oOHz2WxIKGyf3Byt:XxVcpRZWKRDV2 v3QXBFY5VoSGhrAe:orJ0fKfwdrWoYy]
2022/10/20 14:32:44 Work 8 step 6 started
2022/10/20 14:32:44 Work 8 question 6 answer: map[0kk8GfNLjYhS9sE3:cL9FUB42XFNIycO CgJ7e07DCaEBKAHh:test OI84sfR6RhKKYD0F:A4Z2Pvc1y5QGx]
2022/10/20 14:32:44 Work 9 step 6 started
2022/10/20 14:32:44 Work 9 question 6 answer: map[QTceJTBeNOKyQmo4:test Tr1ykYdQXpJ6PhTD:VG6wL4cqiuU08Ad rIggcynir1QXuq1W:SbPZCa1sdaZZAG]
2022/10/20 14:32:44 Work 10 step 6 started
2022/10/20 14:32:45 Work 10 test successfully passed
2022/10/20 14:32:45 Work 1 step 7 started
2022/10/20 14:32:45 Work 1 test successfully passed
2022/10/20 14:32:45 Work 2 step 7 started
2022/10/20 14:32:45 Work 2 question 7 answer: map[5TCPYQiDJzHtyb8J:fhE3CZ3yEBOSeo1x 6dTQmfYpkfBoaAVG:test NvDUxzb1lfNsjmqF:test kDcwMhKAYzel1cHs:0TcTMr4z5usc8z]
2022/10/20 14:32:45 Work 3 step 7 started
2022/10/20 14:32:45 Work 3 question 7 answer: map[A90SWmmlWA9QEr0Y:test FyEwhqkT2s799zNA:test QV9hnqExmLtylrgf:SLpX2n5z22GBq zKbniM4jQxi5G9O1:test]
2022/10/20 14:32:45 Work 4 step 7 started
2022/10/20 14:32:45 Work 4 question 7 answer: map[9t8VEJfui5TnY224:uUcYiQApmv1wGL]
2022/10/20 14:32:45 Work 5 step 7 started
2022/10/20 14:32:46 Work 5 test successfully passed
2022/10/20 14:32:46 Work 6 step 7 started
2022/10/20 14:32:46 Work 6 question 7 answer: map[M3f2ot9BhzYniVMV:6EQ9ceI4dUP2H dnq4vujiDbPu1nS9:test nwtyh5jsQGNiKvc6:RnmqypBrwa2rKO]
2022/10/20 14:32:46 Work 7 step 7 started
2022/10/20 14:32:46 Work 7 question 7 answer: map[5DoTPh9wKnIQmj02:oUdB4E8OLb9ZVSk GU7AQLirzZJPbyMn:WOw30UHls0actC3 mhzWNP3l7PQ9gTZX:IbFFFHsfRIdWe oX9g8cM5tCVBk5en:HSCZxXieq3IWH8]
2022/10/20 14:32:46 Work 8 step 7 started
2022/10/20 14:32:46 Work 8 question 7 answer: map[VrLnKBsYsu6VDK3m:upsFNmIbFlq8Z0on]
2022/10/20 14:32:46 Work 9 step 7 started
2022/10/20 14:32:46 Work 9 test successfully passed
2022/10/20 14:32:46 Work 2 step 8 started
2022/10/20 14:32:47 Work 2 test successfully passed
2022/10/20 14:32:47 Work 3 step 8 started
2022/10/20 14:32:47 Work 3 question 8 answer: map[4UFwWvewLD7zYKHz:xVoPWPlXKGyZI9Od 5bFKrjW587qAB2s8:test VCRmeH5HEkkBKdHi:test dxqXjnIjgStizEkd:0Pk1k14Fve4wC]
2022/10/20 14:32:47 Work 4 step 8 started
2022/10/20 14:32:47 Work 4 test successfully passed
2022/10/20 14:32:47 Work 6 step 8 started
2022/10/20 14:32:47 Work 6 question 8 answer: map[0dq1CV1gE9uaWrgV:YU8gfClFIXhdUn9e 5dfpGQQgkm4B4siJ:lTomAVyfmaeD8 C7NsP7l9omiuAJKF:test K42xyQUs2CnjhF9Z:dNHihee0ikEJjeZ]
2022/10/20 14:32:47 Work 7 step 8 started
2022/10/20 14:32:47 Work 7 test successfully passed
2022/10/20 14:32:47 Work 8 step 8 started
2022/10/20 14:32:48 Work 8 test successfully passed
2022/10/20 14:32:48 Work 3 step 9 started
2022/10/20 14:32:48 Work 3 question 9 answer: map[1BDgsjiZJnd8OvSk:8jKW4LsC9bNrf 1LLyuIqrOAS1mlae:test qyttPRUTiX8AeusJ:ccag79TH5f6NVif wUocNnRZiiDmqxo5:ZKPpdlVACCXof9u]
2022/10/20 14:32:48 Work 6 step 9 started
2022/10/20 14:32:48 Work 6 question 9 answer: map[JuUHqcGAwjIZkjeB:test TOlRU72itxgwFpID:pPJNCezSxMeDlM bQCwrCByoGqPDvc5:3dPzAsxmg3a59u]
2022/10/20 14:32:48 Work 3 step 10 started
2022/10/20 14:32:48 Work 3 question 10 answer: map[cFbDq9YMptP1UdlM:165iL6P2Bz6Xea9 zVev3Rgiotdllif2:MRHCe7UBxdhY2i]
2022/10/20 14:32:48 Work 6 step 10 started
2022/10/20 14:32:48 Work 6 question 10 answer: map[ajGG8TnFxMMbip1c:necmRyID9IKXQOu qwcmrEXWIoZOXir4:O2KwV3Y9etF3iLV wtzq6wJBksldELFf:test]
2022/10/20 14:32:48 Work 3 step 11 started
2022/10/20 14:32:49 Work 3 test successfully passed
2022/10/20 14:32:49 Work 6 step 11 started
2022/10/20 14:32:49 Work 6 test successfully passed
2022/10/20 14:32:49 all works completed
```