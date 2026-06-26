import { useEffect, useState } from "react";
import { useNavigate } from "react-router";

import type { Medicine, MedicineType, MedicineStatus } from "../../api/medicine";
import { getMedicines, medicineTypes, medicineStatuses } from "../../api/medicine";
import DataTable from "../../components/datatable";
import Pagination from "../../components/pagination";
import '../../App.css';

export default function MedicinePage() {
    const [medicines, setMedicines] = useState<Medicine[]>([]);
    const [searchKeyword, setSearchKeyword] = useState("");
    const [filterType, setFilterType] = useState<MedicineType[]>([]);
    const [filterStatus, setFilterStatus] = useState<MedicineStatus[]>([]);
    const [sortBy, setSortBy] = useState("id");
    const [isSortAsc, setIsSortAsc] = useState(true);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [expandedTypeFilter, setExpandedTypeFilter] = useState(true);
    const [expandedStatusFilter, setExpandedStatusFilter] = useState(true);
    const [page, setPage] = useState(1);

    const navigate = useNavigate();

    const toggleTypeFilter = (type: MedicineType) => {
        setFilterType(prev =>
            prev.includes(type) ? prev.filter(t => t !== type) : [...prev, type]
        );
    };
    const toggleStatusFilter = (status: MedicineStatus) => {
        setFilterStatus(prev =>
            prev.includes(status) ? prev.filter(s => s !== status) : [...prev, status]
        );
    };

    const processedMedicines = medicines
        .filter((medicine) =>
            medicine.name.toLowerCase().includes(searchKeyword.toLowerCase())
        )
        .filter((medicine) =>
            filterType.length === 0 || filterType.includes(medicine.type)
        )
        .filter((medicine) =>
            filterStatus.length === 0 || filterStatus.includes(medicine.status)
        )
        .sort((a, b) => {
            let valA: string | number;
            let valB: string | number;
            switch (sortBy) {
                case "name":
                    valA = a.name.toLowerCase();
                    valB = b.name.toLowerCase();
                    break;
                case "strength":
                    valA = a.strengthValue;
                    valB = b.strengthValue;
                    break;
                default:
                    valA = a.id;
                    valB = b.id;
            }

            if (valA < valB) return isSortAsc ? -1 : 1;
            if (valA > valB) return isSortAsc ? 1 : -1;
            return 0;
        });

    const pageSize = 5;

    const totalPages = Math.max(1, Math.ceil(processedMedicines.length / pageSize));
    const safePage = Math.min(page, totalPages);
    const paginatedMedicines = processedMedicines.slice((safePage - 1) * pageSize, safePage * pageSize);


    useEffect(() => {
        async function loadMedicine() {
            setLoading(true);
            setError(null);

            try {
                if (isActive) {
                    const data = await getMedicines();
                    setMedicines(data);
                }

            } catch (err) {
                setError(`Unable to load medicines: ${err}\nPlease try again.`);
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
    }, []);

    if (loading) {
        return (
            <div>
                <p>Loading medicine list...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div>
                <p>{error}</p>
            </div>
        );
    }

    return (
        <div>
            <h2>Medicine</h2>

            <button onClick={() => navigate("/medicine/create")}>Create Medicine</button>

            <input
                value={searchKeyword}
                onChange={(e) => {
                    setSearchKeyword(e.target.value);
                }}
                placeholder="Search medicine"
            />

            <div>
                <label>Sort by:</label>
                <select onChange={(e) => { setSortBy(e.target.value); }}>
                    <option value="id">ID</option>
                    <option value="name">Name</option>
                    <option value="strength">Strength</option>
                </select>
                <button onClick={() => setIsSortAsc(!isSortAsc)}>
                    Sort order: {isSortAsc ? "▲" : "▼"}
                </button>
            </div>

            <div>
                <button onClick={() => setExpandedTypeFilter(!expandedTypeFilter)}>
                    {expandedTypeFilter ? "▼" : "▶"} Filter Type
                </button>
                {expandedTypeFilter && (
                    <div>
                        {medicineTypes.map(type => (
                            <label key={type}>
                                <input
                                    type="checkbox"
                                    checked={filterType.includes(type)}
                                    onChange={() => toggleTypeFilter(type)}
                                />
                                {" "} {type}
                            </label>
                        ))}
                    </div>
                )}
            </div>

            <div>
                <button onClick={() => setExpandedStatusFilter(!expandedStatusFilter)}>
                    {expandedStatusFilter ? "▼" : "▶"} Filter Status
                </button>
                {expandedStatusFilter && (
                    <div>
                        {medicineStatuses.map(status => (
                            <label key={status}>
                                <input
                                    type="checkbox"
                                    checked={filterStatus.includes(status)}
                                    onChange={() => toggleStatusFilter(status)}
                                />
                                {" "} {status}
                            </label>
                        ))}
                    </div>
                )}
            </div>

            <DataTable
                headers={["ID", "Name", "Type", "Strength", "Status", "Actions"]}
                rows={paginatedMedicines.map((medicine) => [
                    medicine.code,
                    medicine.name,
                    medicine.type,
                    `${medicine.strengthValue} ${medicine.strengthUnit}`,
                    medicine.status,
                    <button onClick={() => navigate(`/medicine/${medicine.code}`)}>View Details</button>
                ])}
            />

            <Pagination page={safePage} totalPages={totalPages} onPageChange={setPage} />
        </div>
    );
}