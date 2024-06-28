import type { MasterHealthEnum } from "~/pkg/domain/model/health/masterHealth/master_health_enum.gen";

export type MasterHealth = {
  healthId: number;
  name: string;
  masterHealthEnum: MasterHealthEnum;
};
