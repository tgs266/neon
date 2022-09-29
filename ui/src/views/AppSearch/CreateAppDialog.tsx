import React, { useEffect, useState } from "react";
import {
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Typography,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
} from "@mui/material";
import { CredentialsService } from "../../services/CredentialsService";
import { Credentials } from "../../models/credentials";
import { CreateAppRequest } from "../../models/app";
import { AppService } from "../../services/AppService";

export function CreateAppDialog(props: {
  open: boolean;
  setOpen: (b: boolean) => void;
}) {
  const [credentials, setCredentials] = useState<Credentials[]>([]);
  const [selectedCredentials, setSelectedCredentials] = useState("");
  const [name, setName] = useState("");
  const [repo, setRepo] = useState("");


  useEffect(() => {
    CredentialsService.getAll().then((r) => {
      setCredentials(r.data);
    });
  }, [props.open]);


  const save = () => {
    const req: CreateAppRequest = {
      name,
      credentialName: selectedCredentials,
      repository: repo,
      products: []
    }

    AppService.create(req).then(r => {
      props.setOpen(false)
    })

  }

  return (
    <Dialog maxWidth="md" fullWidth open={props.open} onClose={props.setOpen}>
      <DialogTitle>Create App</DialogTitle>
      <DialogContent>
        {/* <DialogContentText>
          To subscribe to this website, please enter your email address here. We
          will send updates occasionally.
        </DialogContentText> */}
        <TextField
          margin="dense"
          id="name"
          label="Name"
          fullWidth
          variant="standard"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
          <TextField
            margin="dense"
            id="repo"
            label="Repository"
            fullWidth
            variant="standard"
            value={repo}
            onChange={(e) => setRepo(e.target.value)}
          />
          <FormControl margin="dense" variant="standard" fullWidth>
            <InputLabel id="creds">Credentials</InputLabel>
            <Select
              labelId="creds"
              value={selectedCredentials}
              label="Credentials"
              onChange={(e) => setSelectedCredentials(e.target.value)}
            >
              {credentials.map((c) => (
                <MenuItem key={c.name} value={c.name}>{c.name}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </div>
      </DialogContent>
      <DialogActions>
        <Button onClick={() => props.setOpen(false)}>Cancel</Button>
        <Button onClick={() => save()}>Create</Button>
      </DialogActions>
    </Dialog>
  );
}
