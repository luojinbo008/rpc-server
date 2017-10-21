namespace go thriftRPC.guide

struct Guide {
    1:  i16         app_id
    2:  i32         guide_id
    3:  string      nick_name
    4:  string      real_name
    5:  string      id_card
    6:  i64         phone
    7:  i32         birthday
    8:  string      memo
}

service guideThrift {
     bool Register(1:Guide guide),
}