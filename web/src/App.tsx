import { useEffect, useState } from "react";
import type { Medicine, CreateMedicineRequest } from "./api/medicine";
import { createMedicine, deleteMedicine, getMedicines } from "./api/medicine";
import DataTable from "./components/datatable";
import type { SimpleRow } from "./components/datatable";
import './App.css';

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
        <div style={{ padding: 24, marginTop: 40 }}>
            <h1 style={{ fontSize: 24, fontWeight: "bold" }}>Medicine</h1>

            <input
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="New medicine name"
                style={{
                    borderWidth: 1,
                    padding: 8,
                    margin: 12,
                }}
            />

            <button onClick={handleCreate}>Create Medicine</button>

            <DataTable
                headers={["ID", "Name", "Type", "Strength", "Status", "Actions"]}
                rows={medicines.map((medicine) => [
                    medicine.code,
                    medicine.name,
                    medicine.type,
                    `${medicine.strengthValue} ${medicine.strengthUnit}`,
                    medicine.status,
                    <button onClick={() => handleDelete(medicine.code)}>Delete</button>
                ])}
            />
        </div>
    );
}

// import { useState, useEffect } from 'react'
// // import { getHealth } from './api/client';
// import './App.css'

// function App() {
//     // const [status, setStatus] = useState("Loading...");

//     // useEffect(() => {
//     //     getHealth().then(setStatus);
//     // }, []);

//     // return (
//     //     <div className="App">
//     //         <h1>Dingus Backend Status: {status}</h1>
//     //     </div>
//     // );


// }

// export default App;