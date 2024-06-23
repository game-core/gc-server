import type { MasterHealthEnum } from "~/pkg/domain/model/health/masterHealth/MasterHealthEnum.gen";

export type MasterHealth = {
  healthId: number;
  name: string;
  masterHealthEnum: MasterHealthEnum;
};
