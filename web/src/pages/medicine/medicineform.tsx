import { useState, useEffect } from "react";
import { Link, useNavigate, useParams } from "react-router";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type { MedicineType, MedicineFormValues } from "../../api/medicine";
import { createMedicine, getMedicine, medicineTypes, medicineSchema, getUnitsFromType, updateMedicine } from "../../api/medicine";

export default function MedicineFormPage() {
    const [loading, setLoading] = useState(true);
    const { code } = useParams<{ code?: string }>();
    const isUpdate = Boolean(code);
    const navigate = useNavigate();

    const defaultType = medicineTypes[0];
    const defaultStrengthUnits = getUnitsFromType(defaultType);

    const {
        register,
        handleSubmit,
        watch,
        reset,
        setValue,
        formState: { errors },
    } = useForm<MedicineFormValues>({
        resolver: zodResolver(medicineSchema),
        defaultValues: {
            name: "",
            type: defaultType,
            strengthValue: 1,
            strengthUnit: defaultStrengthUnits[0],
            description: "",
        },
    });

    const selectedType = watch("type") || defaultType;
    const selectedStrengthUnit = watch("strengthUnit") || defaultStrengthUnits[0];
    const strengthUnits = getUnitsFromType(selectedType);

    useEffect(() => {
        if (!isUpdate) {
            setLoading(false);
            return;
        }

        async function loadMedicine() {
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

                    // set new default values as fields of loaded medicine
                    reset({
                        name: data.name,
                        type: data.type,
                        strengthValue: data.strengthValue,
                        strengthUnit: data.strengthUnit,
                        description: data.description ?? "",
                    });
                }
            } catch (err) {
                console.error(err);
            } finally {
                setLoading(false);
            }
        }

        let isActive = true;
        loadMedicine();
        return () => {
            isActive = false;
        };
    }, [code, isUpdate, navigate, reset]);

    useEffect(() => {
        if (!strengthUnits.includes(selectedStrengthUnit)) {
            setValue("strengthUnit", strengthUnits[0]);
        }
    }, [selectedStrengthUnit, strengthUnits, setValue]);

    const onSubmit = async (values: MedicineFormValues) => {
        const result = code ? await updateMedicine(code, values) : await createMedicine(values);

        if (!result) {
            alert("Failed to save medicine. Please try again.");
            return;
        }

        console.log(result);

        if (typeof result === "string") {
            navigate(`/medicine/${result}`);
        }
        else {
            navigate(`/medicine/${result.code}`);
        }
    };

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
            <form onSubmit={handleSubmit(onSubmit)}>
                <label htmlFor="name">Name:</label>
                <input id="name" {...register("name")} placeholder="Medicine name here" />
                {errors.name && <div style={{ color: "red" }}>{errors.name.message}</div>}
                <br />

                <label htmlFor="type">Type:</label>
                <select
                    id="type"
                    {...register("type", {
                        onChange: (e) => {
                            const newType = e.target.value as MedicineType;
                            setValue("strengthUnit", getUnitsFromType(newType)[0]);
                        },
                    })}
                >
                    {medicineTypes.map((t) => (
                        <option key={t} value={t}>
                            {t}
                        </option>
                    ))}
                </select>
                {errors.type && <div style={{ color: "red" }}>{errors.type.message}</div>}
                <br />

                <label htmlFor="strengthValue">Strength:</label>
                <input
                    id="strengthValue"
                    type="number"
                    {...register("strengthValue", {
                        valueAsNumber: true,
                    })}
                    placeholder="Strength value"
                />
                <select id="strengthUnit" {...register("strengthUnit")}>
                    {strengthUnits.map((t) => (
                        <option key={t} value={t}>
                            {t}
                        </option>
                    ))}
                </select>
                {errors.strengthValue && <div style={{ color: "red" }}>{errors.strengthValue.message}</div>}
                <br />

                <label htmlFor="description">Description:</label>
                <textarea id="description" {...register("description")} placeholder="Description" />
                {errors.description && <div style={{ color: "red" }}>{errors.description.message}</div>}
                <br />

                <button type="submit">{code ? "Update Medicine" : "Create Medicine"}</button>
            </form>
        </>
    );
}