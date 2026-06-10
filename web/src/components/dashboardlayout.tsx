import { Link, Outlet } from "react-router";

export default function DashboardLayout() {
    return (
        <>
            <div className="header">
                <h1>BODHISOFT</h1>
            </div>
            <div className="dashboard-layout">
                <aside className="sidebar">
                    <nav>
                        <ul style={{ listStyle: "none", padding: 0 }}>
                            <li>
                                <Link to="/medicine">Medicine</Link>
                            </li>
                        </ul>
                    </nav>
                </aside>
                <div className="content">
                    <Outlet />
                </div>
            </div>
        </>
    );
}