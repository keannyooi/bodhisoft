import { useState, useEffect } from "react";
import { Link, useNavigate, useParams, useLocation } from "react-router";
import type { CreateMedicineRequest, MedicineStrengthUnit, MedicineType } from "../../api/medicine";
import { createMedicine, getMedicine, medicineTypes, getUnitsFromType, updateMedicine } from "../../api/medicine";

type MedicineFormErrors = {
    name?: string;
    type?: string;
    strengthValue?: string;
    description?: string;
};

export default function MedicineFormPage() {
    const [name, setName] = useState("");
    const [type, setType] = useState<MedicineType>(medicineTypes[0]);
    const [strengthValue, setStrengthValue] = useState(1);
    const [strengthUnits, setStrengthUnits] = useState<MedicineStrengthUnit[]>(getUnitsFromType(medicineTypes[0]));
    const [strengthUnitIndex, setStrengthUnitIndex] = useState(0);
    const [description, setDescription] = useState("");
    const [errors, setErrors] = useState<MedicineFormErrors>({});
    const [loading, setLoading] = useState(true);

    const { code } = useParams<{ code?: string }>();
    const isUpdate = Boolean(code);
    const navigate = useNavigate();
    const location = useLocation();

    function handleTypeChange(type: string) {
        const medicineType = type as MedicineType;
        setType(medicineType);

        setStrengthUnits(getUnitsFromType(type));
    }

    useEffect(() => {
        if (!isUpdate) {
            setLoading(false);
            return;
        }

        async function loadMedicine() {
            console.log(location);
            if (!code) {
                navigate("/medicine", { replace: true });
                return;
            }

            setLoading(true);

            try {
                if (isActive) {
                    const data = await getMedicine(code);
                    if (!data) {
                        // redirect back to medicine list if no medicine code is given to update action
                        navigate("/medicine", { replace: true });
                        return;
                    }

                    setName(data.name);
                    setStrengthValue(data.strengthValue);
                    setType(data.type);
                    setStrengthUnits(getUnitsFromType(data.type));
                    if (data.description) setDescription(data.description);
                }
            } catch (err) {
                console.log(err);
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

    // TODO: handle race conditions
    async function handleSubmit() {
        const obtainedErrors: MedicineFormErrors = {};
        if (name.trim() === "") {
            obtainedErrors.name = "Name is required";
        }
        if (strengthValue <= 0 || strengthValue > 999) {
            obtainedErrors.strengthValue = "Strength value must be between 1 and 999";
        }


        if (Object.keys(obtainedErrors).length > 0) return setErrors(obtainedErrors);

        const req: CreateMedicineRequest = {
            name: name.trim(),
            type: type as MedicineType,
            strengthValue,
            strengthUnit: strengthUnits[strengthUnitIndex] as MedicineStrengthUnit,
        }
        if (isUpdate || description.trim() !== "") {
            req.description = description.trim();
        }

        let result;
        if (code) {
            result = await updateMedicine(code, req);
        } else {
            result = await createMedicine(req);
        }

        if (!result) {
            return alert("Failed to save medicine. Please try again.");
        }

        return navigate(`/medicine/${result.code}`);
    }

    if (loading) {
        return (
            <div>
                <p>Loading medicine form...</p>
            </div>
        );
    }

    return (
        <>
            <div>
                <h2>{code ? "Update" : "Create"} Medicine</h2>
                <Link to="/medicine">Back to List</Link>
            </div>
            <form onSubmit={(e) => {
                e.preventDefault();
                handleSubmit();
            }}>
                <label htmlFor="name">Name:</label>
                <input
                    id="name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="New medicine name"
                />
                {errors.name && <div style={{ color: "red" }}>{errors.name}</div>}
                <br />

                <label htmlFor="type">Type:</label>
                <select
                    id="type"
                    value={type}
                    onChange={(e) => handleTypeChange(e.target.value)}
                >
                    {medicineTypes.map((t) => (
                        <option key={t} value={t}>{t}</option>
                    ))}
                </select>
                {errors.type && <div style={{ color: "red" }}>{errors.type}</div>}
                <br />

                <label htmlFor="strengthValue">Strength:</label>
                <input
                    id="strengthValue"
                    type="number"
                    value={strengthValue}
                    min={1}
                    max={999}
                    onChange={(e) => setStrengthValue(Number(e.target.value))}
                    placeholder="Strength value"
                />
                <select
                    id="strengthUnit"
                    value={strengthUnitIndex}
                    onChange={(e) => setStrengthUnitIndex(parseInt(e.target.value, 10))}
                >
                    {strengthUnits.map((t, index) => (
                        <option key={t} value={index}>{t}</option>
                    ))}
                </select>
                {
                    errors.strengthValue && <div style={{ color: "red" }}>{errors.strengthValue}</div>
                }
                <br />

                <label htmlFor="description">Description:</label>
                <textarea
                    id="description"
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                    placeholder="Description"
                />
                {
                    errors.description && <div style={{ color: "red" }}>{errors.description}</div>
                }
                <br />

                <button type="submit">{code ? "Update Medicine" : "Create Medicine"}</button>
            </form>
        </>
    )
}