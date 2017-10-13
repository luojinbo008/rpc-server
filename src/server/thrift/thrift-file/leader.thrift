namespace go leader

struct Leader {
    1:  i16         app_id
    2:  i32         leader_id
    3:  string      nick_name
    4:  string      real_name
    5:  string      id_card
    6:  i64         phone
    7:  i32         birthday
    8:  string      memo
}


service leaderThrift {
     bool Register(1:Leader leader),
}