import { Card, Button } from "@mui/material";
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

        <AddCredentialDialog open={open} setOpen={setOpen} />
      </Accordion>
    </Box>
  );
}
