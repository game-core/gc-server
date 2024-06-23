import type { AdminHealth } from "~/pkg/domain/model/health/adminHealth/AdminHealth.gen";
import type { CommonHealth } from "~/pkg/domain/model/health/commonHealth/CommonHealth.gen";
import type { MasterHealth } from "~/pkg/domain/model/health/masterHealth/MasterHealth.gen";

export type HealthCheckResponse = {
  adminHealth: AdminHealth;
  commonHealth: CommonHealth;
  masterHealth: MasterHealth;
};
