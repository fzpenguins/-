namespace go picture

struct Image {
    1: i64 pid
    2: string url
}


struct InsertRequest {
    1: string url
}

struct InsertResponse {
    1: Image image,
    2: BaseResp base;
}

struct SearchByImageRequest {
    1: string url 
}


struct SearchResponse {
    1: list<Image> images;
    2: BaseResp base;
}

struct BaseResp{
    1: i64 code,
    2: string msg,
}

service PictureService {
    InsertResponse Insert(1: InsertRequest req);
    SearchResponse SearchByImage(1: SearchByImageRequest req);
}