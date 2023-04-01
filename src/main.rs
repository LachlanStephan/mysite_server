
use axum::{
    response::Json,
    routing::get,
    Router,
    http::Method
};
use tower_http::cors::{Any, CorsLayer};
use serde_json::{Value, json};


#[tokio::main]
async fn main() {
    let cors = CorsLayer::new() 
        .allow_methods([Method::GET, Method::POST])
        .allow_origin(Any);

    let app = Router::new()
        .route("/", get(plain_text))
        .route("/blog", get(blog))
        .layer(cors);

    axum::Server::bind(&"0.0.0.0:8080".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn home() -> String {
    String::from("Home page")
}

async fn plain_text() -> &'static str {
    "hello there"
}

async fn blog() -> Json<Value> {
    Json(json!({ "data": "blog here" }))
}