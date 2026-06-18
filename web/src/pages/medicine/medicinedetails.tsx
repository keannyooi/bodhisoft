import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router";
import type { Medicine } from "../../api/medicine";
import { deleteMedicine, getMedicine } from "../../api/medicine";

export default function MedicineDetailsPage() {
    const [medicine, setMedicine] = useState<Medicine | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const { code } = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        async function loadMedicine() {
            if (!code) {
                navigate("/medicine", { replace: true });
                return;
            }

            setLoading(true);
            setError(null);

            try {
                if (isActive) {
                    const data = await getMedicine(code);
                    if (!data) {
                        navigate("/medicine", { replace: true });
                        return;
                    }

                    setMedicine(data);
                }
            } catch (err) {
                setError(`Unable to load medicine details: ${err}\nPlease try again.`);
            } finally {
                setLoading(false);
            }
        }

        // boolean race condition handling
        let isActive = true;
        loadMedicine();
        return () => {
            isActive = false
        };
    }, [code, navigate]);

    async function handleDelete() {
        if (!medicine) return;

        await deleteMedicine(medicine.code);
        navigate("/medicine", { replace: true });
    }

    if (loading) {
        return (
            <div>
                <p>Loading medicine details...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div>
                <p>{error}</p>
                <Link to="/medicine">Back to medicine list</Link>
            </div>
        );
    }

    if (!medicine) {
        return null;
    }

    return (
        <>
            <div>
                <h2>Medicine Details: {medicine.code}</h2>
                <Link to="/medicine">Back to List</Link>
            </div>
            <div>
                <div>
                    <table>
                        <tbody>
                            <tr>
                                <td>Name</td>
                                <td>{medicine.name}</td>
                            </tr>
                            <tr>
                                <td>Type</td>
                                <td>{medicine.type}</td>
                            </tr>
                            <tr>
                                <td>Strength</td>
                                <td>
                                    {medicine.strengthValue} {medicine.strengthUnit}
                                </td>
                            </tr>
                            <tr>
                                <td>Description</td>
                                <td>{medicine.description ?? "-"}</td>
                            </tr>
                            <tr>
                                <td>Status</td>
                                <td>{medicine.status}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div>
                    <h3>Administrator Actions</h3>
                    <button onClick={() => navigate(`/medicine/update/${medicine.code}`)}>
                        Update
                    </button>
                    <button onClick={handleDelete}>Delete</button>
                </div>
            </div>
        </>
    );
}