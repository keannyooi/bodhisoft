import { useEffect, useState } from "react";
import type { Medicine, MedicineType, MedicineStatus } from "../../api/medicine";
import { getMedicines, medicineTypes, medicineStatuses } from "../../api/medicine";
import { useNavigate } from "react-router";
import DataTable from "../../components/datatable";
import '../../App.css';

export default function MedicinePage() {
    const [medicines, setMedicines] = useState<Medicine[]>([]);
    const [searchKeyword, setSearchKeyword] = useState("");
    const [filterType, setFilterType] = useState<MedicineType[]>([]);
    const [filterStatus, setFilterStatus] = useState<MedicineStatus[]>([]);
    const [sortBy, setSortBy] = useState("id");
    const [isSortAsc, setIsSortAsc] = useState(true);
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

    async function loadMedicines() {
        const data = await getMedicines();
        setMedicines(data);
    }

    // async function handleCreate() {
    //     // TODO: client-side validation
    //     if (!name.trim()) return;

    //     const req: CreateMedicineRequest = {
    //         name: name,
    //         type: "Tablet",
    //         strengthValue: 50,
    //         strengthUnit: "mg"
    //     };
    //     // if (description.trim() != "") {
    //     //     req.description = description;
    //     // }

    //     await createMedicine(req);
    //     setName("");
    //     await loadMedicines();
    // }

    // async function handleDelete(id: string) {
    //     await deleteMedicine(id);
    //     await loadMedicines();
    // }

    useEffect(() => {
        loadMedicines();
    }, []);

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
                    Sort order: {isSortAsc ? "^" : "v"}
                </button>
            </div>

            <label>Filter Type:</label>
            {
                medicineTypes.map(type => (
                    <label key={type}>
                        <input
                            type="checkbox"
                            checked={filterType.includes(type)}
                            onChange={() => toggleTypeFilter(type)}
                        />
                        {" "} {type}
                    </label>
                ))
            }

            <label>Filter Status:</label>
            {
                medicineStatuses.map(status => (
                    <label key={status}>
                        <input
                            type="checkbox"
                            checked={filterStatus.includes(status)}
                            onChange={() => toggleStatusFilter(status)}
                        />
                        {" "} {status}
                    </label>
                ))
            }

            <DataTable
                headers={["ID", "Name", "Type", "Strength", "Status", "Actions"]}
                rows={processedMedicines.map((medicine) => [
                    medicine.code,
                    medicine.name,
                    medicine.type,
                    `${medicine.strengthValue} ${medicine.strengthUnit}`,
                    medicine.status,
                    <button onClick={() => navigate(`/medicine/${medicine.code}`)}>View Details</button>
                ])}
            />
        </div>
    );
}