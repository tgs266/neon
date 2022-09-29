import { Card, Button, Table, TableHead, TableBody, TableCell, TableRow } from "@mui/material";
import { Box } from "@mui/system";
import React, { useEffect, useState } from "react";
import Accordion from "../../components/Accordion";
import TitleCard from "../../components/DetailsTitleCard";
import { AddCredentialDialog } from "./AddCredentialDialog";

export function Settings() {
  const [open, setOpen] = useState(false)

  return (
    <Box>
      <Accordion title="Credentials">
        <Button onClick={() => setOpen(true)}>Add New</Button>
        {/* <Table sx={{ minWidth: 650 }}>
        <TableHead>
            <TableRow>
                <TableCell>App</TableCell>
                <TableCell align="right">Version</TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {product?.installs && product.installs.map((row) => (
                <TableRow
                    key={row.appName}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">
                        <div style={{ display: "flex", alignItems: "center" }}>
                            <Link to={`/apps/${row.appName}`}>
                                {row.appName}
                            </Link>
                            {row.error &&
                                <Tooltip title={row.error}>
                                    <ErrorIcon color="secondary" />
                                </Tooltip>
                            }
                        </div>
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                        {row.releaseVersion}
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table> */}
        <AddCredentialDialog open={open} setOpen={setOpen} />
      </Accordion>
    </Box>
  );
}
