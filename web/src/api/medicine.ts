const API_URL = "http://192.168.56.1:1337/api/v1";

export const medicineTypes = ["Tablet", "Capsule", "Syrup"] as const;
export type MedicineType = typeof medicineTypes[number];

export const medicineStrengthUnits = ["mg", "g", "mg/ml"] as const;
export type MedicineStrengthUnit = typeof medicineStrengthUnits[number];

export const medicineStatuses = ["Available", "Discontinued"] as const;
export type MedicineStatus = typeof medicineStatuses[number];

export type Medicine = {
    id: number;
    code: string;
    name: string;
    type: MedicineType;
    strengthValue: number;
    strengthUnit: MedicineStrengthUnit;
    description?: string;
    status: MedicineStatus;
};

// request stuff
export interface CreateMedicineRequest {
    name: string;
    type: MedicineType;
    strengthValue: number;
    strengthUnit: MedicineStrengthUnit;
    description?: string;
}
export interface UpdateMedicineRequest {
    name?: string;
    type?: MedicineType;
    strengthValue?: number;
    strengthUnit?: MedicineStrengthUnit;
    description?: string;
    status?: MedicineStatus
}

export function getUnitsFromType(type: string): Array<MedicineStrengthUnit> {
    const medicineType = type as MedicineType;
    switch (medicineType) {
        case "Tablet":
        case "Capsule":
            return ["mg", "g"];
        case "Syrup":
            return ["mg/ml"];
        default:
            return [];
    }
}

export async function getMedicines(): Promise<Medicine[]> {
    const res = await fetch(`${API_URL}/medicines`);
    return res.json();
}

export async function getMedicine(code: string): Promise<Medicine | null> {
    const res = await fetch(`${API_URL}/medicines/${code}`);
    console.log(res);
    if (!res.ok) {
        return null;
    }
    return res.json();
}

export async function createMedicine(req: CreateMedicineRequest): Promise<Medicine | null> {
    const res = await fetch(`${API_URL}/medicines`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(req),
    });

    if (!res.ok) {
        return null;
    }
    return res.json();
}

export async function updateMedicine(code: string, req: UpdateMedicineRequest): Promise<Medicine | null> {
    console.log(req);
    const res = await fetch(`${API_URL}/medicines/${code}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(req),
    });

    if (!res.ok) {
        return null;
    }
    return res.json();
}

export async function deleteMedicine(code: string): Promise<void> {
    await fetch(`${API_URL}/medicines/${code}`, {
        method: "DELETE",
    });
}