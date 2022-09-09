import React, { useState } from "react"
import { Collapse, IconButton, Table, TableBody, TableCell, TableHead, TableRow, Typography } from '@mui/material';
import { QueuedChange } from "../../models/change";
import { DateTime } from "luxon";
import { Box } from "@mui/system";
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';

function QueuedChangesTableRow(props: { change: QueuedChange }) {
    const { change } = props
    const [open, setOpen] = useState(false)

    return <>
        <TableRow
            key={change.id}
            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
        >
            <TableCell>
                <IconButton
                    disabled={change.details === ""}
                    aria-label="expand row"
                    size="small"
                    onClick={() => setOpen(!open)}
                >
                    {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
                </IconButton>
            </TableCell>
            <TableCell component="th" scope="row">
                {change.release.productName}
            </TableCell>
            <TableCell component="th" scope="row" align="right">
                {change.type}
            </TableCell>
            <TableCell component="th" scope="row" align="right">
                {DateTime.fromISO(change.lastChecked).toRelative()}
            </TableCell>
        </TableRow>
        <TableRow>
            <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
                <Collapse in={open} timeout="auto" unmountOnExit>
                    <Box sx={{ margin: 1 }}>
                        <Typography sx={{whiteSpace: "pre-line"}} variant="body1" component="div">
                            {change.details}
                        </Typography>
                    </Box>
                </Collapse>
            </TableCell>
        </TableRow>
    </>
}

export default function QueuedChangesTable(props: { changes: QueuedChange[] }) {
    const { changes } = props
    return <Table sx={{ minWidth: 650 }}>
        <TableHead>
            <TableRow>
                <TableCell />
                <TableCell>Product</TableCell>
                <TableCell align="right">Type</TableCell>
                <TableCell align="right">Last Checked</TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {changes.map((row) => (
                <QueuedChangesTableRow change={row} key={row.id} />
            ))}
        </TableBody>
    </Table>
}