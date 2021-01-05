mod model;

use actix_web::{get, web, App, HttpServer, Responder};
use  model::safe;

#[get("/{id}/{name}/index.html")]
async fn index(web::Path((id, name)): web::Path<(u32, String)>) -> impl Responder {
    format!("Hello {}! id:{}", name, id)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(index))
        .bind("127.0.0.1:9080")?
        .run()
        .await
}

fn ex1() -> () {
    safe::test()
}
