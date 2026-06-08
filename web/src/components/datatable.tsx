import React from 'react';

export type SimpleRow = Array<string | number | React.ReactNode>;

export default function DataTable({ headers, rows }: {
    headers: string[];
    rows: SimpleRow[];
}) {
    if (rows.length === 0) {
        return <div>No data</div>;
    }

    return (
        <table>
            <thead>
                <tr>
                    {
                        headers.map((h) => (
                            <th key={h} > {h} </th>
                        ))
                    }
                </tr>
            </thead>
            <tbody>
                {
                    rows.map((row, i) => {
                        // console.log(row, i);
                        return (
                            <tr key={i} >
                                {

                                    row.map((cell, j) => {
                                        if (cell === null || cell === undefined) {
                                            return <td key={j} >-</td>;
                                        }
                                        else {
                                            return <td key={j} > {cell} </td>;
                                        }
                                    })
                                }
                            </tr>
                        );
                    })
                }
            </tbody>
        </table>
    );
}