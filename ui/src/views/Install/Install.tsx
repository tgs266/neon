import React, { useEffect, useState } from "react";
import Grid from "@mui/material/Grid";
import { Card, IconButton, Paper, Tab, Tabs, Typography } from "@mui/material";
import { AppService } from "../../services/AppService";
import { Link, useSearchParams } from "react-router-dom";
import { useParams } from "react-router";
import { Box } from "@mui/system";
import TitleCard from "../../components/DetailsTitleCard";
import Statistic from "../../components/Statistic";
import { Install } from "../../models/install";
import { ResourceList } from "../../models/resources";
import ResourceContainer from "../../components/Resources/ResourceContainer";
import ConfigEditor from "../../components/Editor/ConfigEditor";
import CodeIcon from "@mui/icons-material/Code";
import DifferenceIcon from "@mui/icons-material/Difference";

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    "aria-controls": `simple-tabpanel-${index}`,
  };
}

export function Install() {
  const [install, setInstall] = useState<Install>(null);
  const [resources, setResources] = useState<ResourceList>(null);
  const { name, productName } = useParams();
  const [search, setSearch] = useSearchParams();

  useEffect(() => {
    AppService.getInstall(name, productName).then((r) => {
      setInstall(r.data);
    });
    AppService.getInstallResources(name, productName).then((r) => {
      setResources(r.data);
    });
  }, [name]);

  if (search.get("tab") == null) {
    setSearch({ tab: "0" }, { replace: true });
  }

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setSearch({ tab: newValue.toString() }, { replace: true });
  };

  return (
    <div
      style={{ display: "flex", flexDirection: "column", height: "calc(100%)" }}
    >
      <TitleCard
        title={
          <div style={{ display: "flex", alignItems: "center" }}>
            <div>
              {name}/{productName}
            </div>
          </div>
        }
        sx={{ flexShrink: 0, zIndex: 100, mb: 0 }}
      >
        <Box sx={{ mt: 1, display: "flex", alignItems: "center", gap: 2 }}>
          <Statistic
            label="Created At"
            value={new Date(install?.createdAt).toLocaleString()}
          />
          <Statistic
            label="Updated At"
            value={new Date(install?.updatedAt).toLocaleString()}
          />
        </Box>
      </TitleCard>

      {install?.error && (
        <Paper sx={{ mt: 1, p: 1, backgroundColor: "rgba(244, 67, 54, .1)" }}>
          <Typography
            sx={{ whiteSpace: "pre-line" }}
            variant="body1"
            component="div"
          >
            {install.error}
          </Typography>
        </Paper>
      )}

      <Box sx={{}}>
        <Tabs
          sx={{ ml: 0 }}
          value={parseInt(search.get("tab"))}
          onChange={handleChange}
          aria-label="basic tabs example"
        >
          <Tab label="Resources" {...a11yProps(0)} />
          <Tab label="Config" {...a11yProps(1)} />
        </Tabs>
      </Box>

      {search.get("tab") === "0" && (
        <Box sx={{ flexGrow: 1, height: "100%", overflowY: "auto" }}>
          <ResourceContainer appName={name} resources={resources} />
        </Box>
      )}
      {search.get("tab") === "1" && (
        <>
          <Card
            sx={{ mt: 1, flexGrow: 1, maxHeight: "500px", overflowY: "auto", display: "flex", alignItems: "flex-start" }}
          >
            <div style={{display: "flex", flexDirection: "column", gap: "4px", marginLeft: "8px"}}>
            <IconButton color="primary">
              <CodeIcon />
            </IconButton>
            <IconButton color="secondary">
              <DifferenceIcon />
            </IconButton>
            </div>
            <div style={{marginTop: "8px", height: "calc(100% - 8px)", width: "100%"}}>
            <ConfigEditor />
            </div>
          </Card>
        </>
      )}
    </div>
  );
}
