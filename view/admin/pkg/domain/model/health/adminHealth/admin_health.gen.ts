import type { AdminHealthEnum } from "~/pkg/domain/model/health/adminHealth/admin_health_enum.gen";

export type AdminHealth = {
  healthId: number;
  name: string;
  adminHealthEnum: AdminHealthEnum;
};
