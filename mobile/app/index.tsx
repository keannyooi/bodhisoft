import { Text, View } from "react-native";
import { getHealth } from "../src/api/client";
import { useEffect, useState } from "react";

export default function Index() {
    const [status, setStatus] = useState("Loading...");

    useEffect(() => {
        getHealth().then(setStatus);
    }, []);

    return (
        <View
            style={{
                flex: 1,
                justifyContent: "center",
                alignItems: "center",
            }}
        >
            <Text>Backend status: {status}</Text>
        </View>
    );
}
