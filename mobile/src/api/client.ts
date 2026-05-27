export const API_URL = "http://192.168.56.1:1337";

export async function getHealth() {
    const res = await fetch(API_URL + "/health");
    return res.text();
}