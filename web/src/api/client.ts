export const API_URL = "http://localhost:1337";

export async function getHealth() {
    const res = await fetch(API_URL + "/health");
    return res.text();
}