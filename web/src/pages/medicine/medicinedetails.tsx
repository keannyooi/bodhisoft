import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router";

import type { Medicine } from "../../api/medicine";
import { deleteMedicine, getMedicine } from "../../api/medicine";
import DataTable from "../../components/datatable";
import '../../App.css';

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
        <div className="page-shell">
            <div className="page-header">
                <h2>Medicine Details</h2>
                <Link to="/medicine">Back to List</Link>
            </div>

            <div className="details-group">
                <div className="section-card">
                    <DataTable
                        headers={["ID", medicine.code]}
                        rows={[
                            ["Name", medicine.name],
                            ["Type", medicine.type],
                            ["Strength", `${medicine.strengthValue} ${medicine.strengthUnit}`],
                            ["Description", medicine.description ?? "-"],
                            ["Status", medicine.status],
                        ]}
                    />
                </div>

                <div className="section-card">
                    <h3>Administrator Actions</h3>
                    <div className="action-stack">
                        <button onClick={() => navigate(`/medicine/update/${medicine.code}`)}>
                            Update
                        </button>
                        {/* TODO: make this a Discontinue button if the medicine already sees use */}
                        <button onClick={handleDelete}>Delete</button>
                    </div>
                </div>
            </div>
        </div>
    );
}