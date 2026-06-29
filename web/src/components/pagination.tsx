export default function Pagination({
    page,
    totalPages,
    onPageChange,
}: {
    page: number;
    totalPages: number;
    onPageChange: React.Dispatch<React.SetStateAction<number>>;
}) {
    return (
        <div className="pagination">
            <button onClick={() => onPageChange((p) => Math.max(p - 1, 1))}>
                Previous
            </button>
            <span>
                Page {page} of {totalPages}
            </span>
            <button onClick={() => onPageChange((p) => Math.min(p + 1, totalPages))}>
                Next
            </button>
        </div>
    );
}

export const pageSize = 10;
