import { Routes, Route } from "react-router";
import MedicinePage from "./pages/medicine/medicine";
import MedicineDetailsPage from "./pages/medicine/medicinedetails";
import MedicineCreatePage from "./pages/medicine/medicinecreate";
import NotFoundPage from "./pages/notfound";
import './App.css';


export default function App() {
    return (
        <Routes>
            {/* Medicine */}
            <Route path="/medicine" element={<MedicinePage />}></Route>
            <Route path="/medicine/:code" element={<MedicineDetailsPage />}></Route>
            <Route path="/medicine/create" element={<MedicineCreatePage />}></Route>

            {/* Catch-all 404 page */}
            <Route path="*" element={<NotFoundPage />} />
        </Routes>
    );
}