import React, { useEffect, useState } from "react";
import {
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
} from "@mui/material";
import { AppService } from "../../services/AppService";

export function CommitDialog(props: {
  open: boolean;
  setOpen: (b: boolean) => void;
  appName: string;
  productName: string;
  data: string;
}) {
  const [msg, setMsg] = useState("Update config");

  const save = () => {
    AppService.updateInstallConfig(props.appName, props.productName, {data: props.data, message: msg}).then(r => {
      props.setOpen(false)
    })

  }

  return (
    <Dialog maxWidth="sm" fullWidth open={props.open} onClose={props.setOpen}>
      <DialogTitle>Commit</DialogTitle>
      <DialogContent>
        <TextField
          margin="dense"
          id="name"
          label="Commit Message"
          fullWidth
          variant="standard"
          value={msg}
          helperText="Commit message cannot be blank"
          error={msg === ""}
          onChange={(e) => setMsg(e.target.value)}
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={() => props.setOpen(false)}>Cancel</Button>
        <Button onClick={() => save()} disabled={msg === ""}>Commit</Button>
      </DialogActions>
    </Dialog>
  );
}
