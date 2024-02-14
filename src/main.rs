use reqwest::header::ACCEPT;
use serde::Deserialize;

#[derive(Deserialize, Debug)]
struct IpifyResponse {
    ip: String,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = reqwest::Client::new();
    let response = client.get("https://api.ipify.org?format=json")
        .header(ACCEPT, "application/json").send().await?;

    match response.status() {
        reqwest::StatusCode::OK => {
            match response.json::<IpifyResponse>().await {
                Ok(parsed) => {
                    println!("Our IP is {}", parsed.ip)
                },
                Err(_) => {
                    println!("JSON Deserialize failed")
                }
            }
        },
        _ => {
            println!("Request failed");
        }
    }


    // let result = response.text().await?;
    // print!("Result: {:#?}", result);



    Ok(())
}
