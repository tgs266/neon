import { Card, Button, Table, TableHead, TableBody, TableCell, TableRow } from "@mui/material";
import { Box } from "@mui/system";
import React, { useEffect, useState } from "react";
import Accordion from "../../components/Accordion";
import TitleCard from "../../components/DetailsTitleCard";
import { Credentials } from "../../models/credentials";
import { CredentialsService } from "../../services/CredentialsService";
import { AddCredentialDialog } from "./AddCredentialDialog";
import CheckIcon from '@mui/icons-material/Check';
import CloseIcon from '@mui/icons-material/Close';

export function Settings() {
  const [open, setOpen] = useState(false)
  const [credentials, setCredentials] = useState<Credentials[]>([])
  useEffect(() => {
    CredentialsService.getAll().then(r => {
      setCredentials(r.data)
    })
  }, [open])
  return (
    <Box>
      <Accordion title="Credentials">
        
        <Table sx={{ minWidth: 650 }}>
        <TableHead>
            <TableRow>
                <TableCell>Name</TableCell>
                <TableCell align="right">Basic Auth</TableCell>
                <TableCell align="right">Token Auth</TableCell>
                <TableCell align="right"><Button onClick={() => setOpen(true)}>Add New</Button></TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {credentials.map((row) => (
                <TableRow
                    key={row.name}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">
                        {row.name}
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                        {row.basicAuth ? <CheckIcon color="success" /> : <CloseIcon color="error" />}
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                        {row.tokenAuth ? <CheckIcon color="success" /> : <CloseIcon color="error" />}
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
        <AddCredentialDialog open={open} setOpen={setOpen} />
      </Accordion>
    </Box>
  );
}
