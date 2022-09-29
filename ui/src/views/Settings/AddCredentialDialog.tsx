import React, { useEffect, useState } from "react";
import {
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  FormControlLabel,
  Switch,
  Typography,
  FormGroup,
} from "@mui/material";
import { CreateCredentialsRequest } from "../../models/credentials";
import { CredentialsService } from "../../services/CredentialsService";

export function AddCredentialDialog(props: {
  open: boolean;
  setOpen: (b: boolean) => void;
}) {

  const [basicAuth, setBasicAuth] = useState(true)
  const [name, setName] = useState("")
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [token, setToken] = useState("")

  const save = () => {
    const req: CreateCredentialsRequest = {
      name
    }
    if (basicAuth) {
      req.username = username
      req.password = password
    } else {
      req.token = token
    }
    CredentialsService.create(req).then(r => {
      props.setOpen(false)
    })
  }

  return (
    <Dialog maxWidth="sm" fullWidth open={props.open} onClose={props.setOpen}>
      <DialogTitle>Add Credential</DialogTitle>
      <DialogContent>
        <div style={{ display: "flex", alignItems: "center" }}>
          <TextField
            margin="dense"
            id="name"
            label="Name"
            variant="standard"
            value={name}
            onChange={(e) => {setName(e.target.value)}}
            error={name.includes(" ")}
            helperText="Name cannot include spaces"
            sx={{ flexGrow: 1 }}
          />
          <FormGroup>
            <FormControlLabel
              sx={{ flexShrink: 0 }}
              control={<Switch defaultChecked value={basicAuth} onChange={() => setBasicAuth(!basicAuth)} />}
              label="Basic Auth"
            />
          </FormGroup>
        </div>
        {basicAuth && <div style={{ display: "flex", alignItems: "center", gap: "8px" }}>
        <TextField
            margin="dense"
            id="name"
            label="Username"
            variant="standard"
            fullWidth
            value={username}
            onChange={(e) => {setUsername(e.target.value)}}
          />
          <TextField
            margin="dense"
            id="name"
            label="Password"
            type="password"
            variant="standard"
            fullWidth
            value={password}
            onChange={(e) => {setPassword(e.target.value)}}
          />
        </div>}
        {!basicAuth && <div style={{ display: "flex", alignItems: "center" }}>
        <TextField
            margin="dense"
            id="name"
            label="Token"
            variant="standard"
            fullWidth
            value={token}
            onChange={(e) => {setToken(e.target.value)}}
            error={token.includes(" ")}
            helperText="Token cannot include spaces"
          />
        </div>}
      </DialogContent>
      <DialogActions>
        <Button onClick={() => props.setOpen(false)}>Cancel</Button>
        <Button onClick={() => save()}>Save</Button>
      </DialogActions>
    </Dialog>
  );
}
