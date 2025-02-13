import React, { useContext, useEffect, useMemo } from "react";
import { RouterStep } from "./steps/RouterStep";
import { ExperimentStep } from "./steps/ExperimentStep";
import { EnricherStep } from "./steps/EnricherStep";
import { EnsemblerStep } from "./steps/EnsemblerStep";
import { OutcomeStep } from "./steps/OutcomeStep";
import schema from "./validation/schema";
import { useTuringApi } from "../../../hooks/useTuringApi";
import { ConfirmationModal } from "../../../components/confirmation_modal/ConfirmationModal";
import { RouterCreationSummary } from "./components/RouterCreationSummary";
import { FormContext, StepsWizardHorizontal, addToast } from "@gojek/mlp-ui";
import ExperimentEngineContext from "../../../providers/experiments/context";
import { useConfig } from "../../../config";

export const CreateRouterForm = ({ projectId, onCancel, onSuccess }) => {
  const {
    appConfig: {
      scaling: { maxAllowedReplica },
    },
  } = useConfig();

  const validationSchema = useMemo(
    () => schema(maxAllowedReplica),
    [maxAllowedReplica]
  );

  const { data: router } = useContext(FormContext);
  const protocol = router.config.protocol;

  const [submissionResponse, submitForm] = useTuringApi(
    `/projects/${projectId}/routers`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
    },
    {},
    false
  );

  const { experimentEngineOptions, getEngineProperties } = useContext(
    ExperimentEngineContext
  );

  useEffect(() => {
    if (submissionResponse.isLoaded && !submissionResponse.error) {
      addToast({
        id: "submit-success-create",
        title: "New Turing router is created!",
        color: "success",
        iconType: "check",
      });
      onSuccess(submissionResponse.data.id);
    }
  }, [submissionResponse, onSuccess]);

  const onSubmit = () => submitForm({ body: JSON.stringify(router) });

  const steps = [
    {
      title: "Router",
      children: <RouterStep projectId={projectId} />,
      validationSchema: validationSchema[0],
      validationContext: { protocol },
    },
    {
      title: "Experiments",
      children: <ExperimentStep projectId={projectId} />,
      validationSchema: validationSchema[1],
      validationContext: { experimentEngineOptions, getEngineProperties },
    },
    {
      title: "Enricher",
      children: <EnricherStep projectId={projectId} />,
      validationSchema: validationSchema[2],
    },
    {
      title: "Ensembler",
      children: <EnsemblerStep projectId={projectId} />,
      validationSchema: validationSchema[3],
    },
    {
      title: "Outcome Tracking",
      children: <OutcomeStep projectId={projectId} />,
      validationSchema: validationSchema[4],
    },
  ];

  return (
    <ConfirmationModal
      title="Deploy Turing Router"
      content={<RouterCreationSummary router={router} />}
      isLoading={submissionResponse.isLoading}
      onConfirm={onSubmit}
      confirmButtonText="Deploy"
      confirmButtonColor="primary"
    >
      {(onSubmit) => (
        <StepsWizardHorizontal
          steps={steps}
          onCancel={onCancel}
          onSubmit={onSubmit}
          submitLabel="Deploy"
        />
      )}
    </ConfirmationModal>
  );
};
