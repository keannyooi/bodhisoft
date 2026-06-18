import { Routes, Route } from "react-router";
import DashboardLayout from "./components/dashboardlayout";
import './App.css';

import MedicinePage from "./pages/medicine/medicine";
import MedicineDetailsPage from "./pages/medicine/medicinedetails";
import MedicineFormPage from "./pages/medicine/medicineform";
import NotFoundPage from "./pages/notfound";

export default function App() {
    return (
        <Routes>
            <Route element={<DashboardLayout />}>
                {/* Medicine */}
                <Route path="/medicine" element={<MedicinePage />}></Route>
                <Route path="/medicine/create" element={<MedicineFormPage />}></Route>
                <Route path="/medicine/update/:code" element={<MedicineFormPage />}></Route>
                <Route path="/medicine/:code" element={<MedicineDetailsPage />}></Route>
            </Route>

            {/* Catch-all 404 page */}
            <Route path="*" element={<NotFoundPage />} />
        </Routes>
    );
}