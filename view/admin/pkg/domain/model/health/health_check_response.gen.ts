import type { AdminHealth } from "~/pkg/domain/model/health/adminHealth/admin_health.gen";
import type { CommonHealth } from "~/pkg/domain/model/health/commonHealth/common_health.gen";
import type { MasterHealth } from "~/pkg/domain/model/health/masterHealth/master_health.gen";

export type HealthCheckResponse = {
  adminHealth: AdminHealth;
  commonHealth: CommonHealth;
  masterHealth: MasterHealth;
};
