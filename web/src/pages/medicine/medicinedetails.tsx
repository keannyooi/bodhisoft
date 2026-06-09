import { useParams } from "react-router";

export default function MedicineDetailsPage() {
    const { code } = useParams();

    return <h1>Medicine Details: {code}</h1>;
}