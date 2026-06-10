import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router";
import type { Medicine } from "../../api/medicine";
import { deleteMedicine, getMedicine } from "../../api/medicine";

export default function MedicineDetailsPage() {
    const [medicine, setMedicine] = useState(null as unknown as Medicine);
    const { code } = useParams();
    const navigate = useNavigate();

    async function loadMedicine() {
        if (!code) return;

        const data = await getMedicine(code);
        if (!data) return navigate("/medicine");

        setMedicine(data);
    }

    async function handleDelete(id: string) {
        await deleteMedicine(id);
        return navigate("/medicine");
    }

    useEffect(() => {
        loadMedicine();
    }, []);

    return (
        <>
            <div>
                <h2>Medicine Details: {code}</h2>
                <Link to="/medicine">Back to List</Link>
            </div>
            <div>
                <div>
                    <table>
                        <tbody>
                            <tr>
                                <td>Name</td>
                                <td>{medicine?.name}</td>
                            </tr>
                            <tr>
                                <td>Type</td>
                                <td>{medicine?.type}</td>
                            </tr>
                            <tr>
                                <td>Strength</td>
                                <td>{medicine?.strengthValue} {medicine?.strengthUnit}</td>
                            </tr>
                            <tr>
                                <td>Description</td>
                                <td>{medicine?.description ?? "-"}</td>
                            </tr>
                            <tr>
                                <td>Status</td>
                                <td>{medicine?.status}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div>
                    <h3>Administrator Actions</h3>
                    <button onClick={() => handleDelete(medicine?.code)}>Delete Medicine</button>
                </div>
            </div>


        </>
    );
}