const API_URL = "http://192.168.56.1:1337/api/v1";

export type Medicine = {
    id: Number;
    code: string;
    name: string;
    type: string;
    strengthValue: Number;
    strengthUnit: string;
    description?: string;
    status: string;
};

// request stuff
export interface CreateMedicineRequest {
    name: string;
    type: string;
    strengthValue: Number;
    strengthUnit: string;
    description?: string;
}
export interface UpdateMedicineRequest {
    name?: string;
    type?: string;
    strengthValue?: Number;
    strengthUnit?: string;
    description?: string;
    status?: string
}

export async function getMedicines(): Promise<Medicine[]> {
    const res = await fetch(`${API_URL}/medicines`);
    return res.json();
}

export async function createMedicine(req: CreateMedicineRequest): Promise<Medicine> {
    const res = await fetch(`${API_URL}/medicines`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(req),
    });

    return res.json();
}

export async function updateMedicine(code: string, req: UpdateMedicineRequest): Promise<Medicine> {
    const res = await fetch(`${API_URL}/medicines/${code}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(req),
    });

    return res.json();
}

export async function deleteMedicine(code: string): Promise<void> {
    await fetch(`${API_URL}/medicines/${code}`, {
        method: "DELETE",
    });
}