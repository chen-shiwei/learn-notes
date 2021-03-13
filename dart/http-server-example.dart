import 'dart:convert';
import 'dart:io';

void main() async {
  HttpServer.bind(InternetAddress.anyIPv4, 8080).then((server) {
    print("server run ${server.address}:${server.port}");
    server.listen((event) {
      // TODO:router
      switch (event.method) {
        case "GET":
          GetPost(event);
          GetPosts(event);
          return;
        case "POST":
          CreatePost(event);
          return;
        case "DELETE":
          DeletePost(event);
          return;
        case "PUT":
          PutPost(event);
          return;
        default:
          responseJson(event, HttpStatus.badRequest, 'not support method');
      }
    });
  });
}

class Post {
  int id;
  String title;
  String content;

  Post(this.id, this.title, this.content) {}

  Map<String, dynamic> toJson() => {
        'id': id,
        'title': title,
        'content': content,
      };
}

var dict = new Map<int, Post>();

Map<String, Function> routers = {
  "GET": GetPost,
  "DELETE": DeletePost,
  "GET": GetPosts,
  "POST": CreatePost,
};

void GetPosts(HttpRequest ctx) {
  if (ctx.uri.toString() != "/posts") {
    return;
  }
  print("GetPosts:${ctx.uri}");
  ctx.response.statusCode = HttpStatus.ok;
  var list = [];
  dict.forEach((key, value) {
    list.add(value);
  });
  responseJsonWithData(ctx, HttpStatus.ok, "success", list);
}

void CreatePost(HttpRequest ctx) async {
  var data = new Map();
  try {
    var content = await utf8.decoder.bind(ctx).join(); /*2*/
    data = jsonDecode(content) as Map; /*3*/
    print("CreatePost:${data}");
  } catch (e) {
    responseJson(ctx, HttpStatus.badRequest, e.toString());
    return;
  }

  nowID += 1;
  var id = nowID;
  var post = new Post(id, data["title"], data["content"]);
  dict[id] = post;
  responseJsonWithData(ctx, HttpStatus.ok, 'success', {"id": id});
}

var nowID = 0;

void GetPost(HttpRequest ctx) {
  if (ctx.uri.toString() == "/posts") {
    return;
  }
  print("GetPost:${ctx.uri}");
  var id = 0;
  try {
    id = int.parse(ctx.uri.queryParameters["id"].toString());
  } catch (e) {
    responseJson(ctx, HttpStatus.badRequest, "id error");
    return;
  }
  var specPost = dict[id];
  if (specPost == null) {
    responseJson(ctx, HttpStatus.notFound, 'post not exists!');
    return;
  }

  responseJsonWithData(ctx, HttpStatus.notFound, 'query success', specPost);
}

void DeletePost(HttpRequest ctx) {
  print("DeletePost:${ctx.uri}");
  var id = int.parse(ctx.uri.queryParameters["id"].toString());
  var specPost = dict[id];
  if (specPost == null) {
    responseJson(ctx, HttpStatus.notFound, 'post not exists!');
    return;
  }
  dict.remove(id);
  responseJsonWithData(ctx, HttpStatus.notFound, 'remove success', specPost);
}

void PutPost(HttpRequest ctx) async {
  var data = new Map();
  try {
    var content = await utf8.decoder.bind(ctx).join(); /*2*/
    data = jsonDecode(content) as Map; /*3*/
    print("PutPost:${data}");
  } catch (e) {
    responseJson(ctx, HttpStatus.badRequest, e.toString());
    return;
  }

  print("put_id:${ctx.uri}");
  var id = int.parse(ctx.uri.queryParameters["id"].toString());
  print("id:${dict[id]}");
  var specPost = dict[id];
  if (specPost == null) {
    responseJson(ctx, HttpStatus.notFound, 'post not exists!');
    return;
  }
  specPost.title = data['title'];
  specPost.content = data['content'];
  dict[id] = specPost;
  responseJsonWithData(ctx, HttpStatus.notFound, 'update success', specPost);
}

void responseJson(HttpRequest ctx, int status, String msg) {
  ctx.response.headers.set("Content-Type", "application/json");
  ctx.response.statusCode = status;
  ctx.response.write(jsonEncode({"code": status, "msg": msg}));
  ctx.response.close();
}

void responseJsonWithData(HttpRequest ctx, int status, String msg, Object ob) {
  ctx.response.headers.set("Content-Type", "application/json");
  ctx.response.statusCode = status;
  ctx.response.write(jsonEncode({"code": status, "msg": msg, "data": ob}));
  ctx.response.close();
}
