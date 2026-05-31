import { useEffect, useState } from "react";
import { Button, FlatList, Text, TextInput, View } from "react-native";
import {
    Medicine,
    CreateMedicineRequest,
    UpdateMedicineRequest,
    createMedicine,
    deleteMedicine,
    getMedicines,
    updateMedicine,
} from "../src/api/medicine";

export default function App() {
    const [medicines, setMedicines] = useState<Medicine[]>([]);
    const [name, setName] = useState("");

    async function loadMedicines() {
        const data = await getMedicines();
        setMedicines(data);
    }

    async function handleCreate() {
        // TODO: client-side validation
        if (!name.trim()) return;

        const req: CreateMedicineRequest = {
            name: name,
            type: "Tablet",
            strengthValue: 50,
            strengthUnit: "mg"
        };
        // if (description.trim() != "") {
        //     req.description = description;
        // }

        await createMedicine(req);
        setName("");
        await loadMedicines();
    }

    // async function handleToggle(Medicine: Medicine) {
    //     await updateMedicine({
    //         ...medicines,
    //         completed: !medicines.completed,
    //     });

    //     await loadMedicines();
    // }

    async function handleDelete(id: string) {
        await deleteMedicine(id);
        await loadMedicines();
    }

    useEffect(() => {
        loadMedicines();
    }, []);

    return (
        <View style={{ padding: 24, marginTop: 40 }}>
            <Text style={{ fontSize: 24, fontWeight: "bold" }}>Medicine</Text>

            <TextInput
                value={name}
                onChangeText={setName}
                placeholder="New medicine name"
                style={{
                    borderWidth: 1,
                    padding: 8,
                    marginVertical: 12,
                }}
            />

            <Button title="Create Medicine" onPress={handleCreate} />

            <FlatList
                data={medicines}
                keyExtractor={(item) => item.code}
                renderItem={({ item }) => (
                    <View style={{ marginTop: 16 }}>
                        <Text>
                            {item.name}
                        </Text>
                        <Button title="Delete" onPress={() => handleDelete(item.code)} />
                    </View>
                )}
            />
        </View>
    );
}