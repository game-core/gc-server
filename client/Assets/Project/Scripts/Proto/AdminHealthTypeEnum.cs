// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: health/adminHealth/admin_health_type_enum.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Api.Game {

  /// <summary>Holder for reflection information generated from health/adminHealth/admin_health_type_enum.proto</summary>
  public static partial class AdminHealthTypeEnumReflection {

    #region Descriptor
    /// <summary>File descriptor for health/adminHealth/admin_health_type_enum.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AdminHealthTypeEnumReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Ci9oZWFsdGgvYWRtaW5IZWFsdGgvYWRtaW5faGVhbHRoX3R5cGVfZW51bS5w",
            "cm90bxIIYXBpLmdhbWUqMgoPQWRtaW5IZWFsdGhUeXBlEg0KCUFkbWluTm9u",
            "ZRAAEhAKDEFkbWluU3VjY2VzcxABQrgBCgxjb20uYXBpLmdhbWVCGEFkbWlu",
            "SGVhbHRoVHlwZUVudW1Qcm90b1ABWk1naXRodWIuY29tL2dhbWUtY29yZS9n",
            "Yy1zZXJ2ZXIvYXBpL2dhbWUvcHJlc2VudGF0aW9uL3Byb3RvL2hlYWx0aC9h",
            "ZG1pbkhlYWx0aKICA0FHWKoCCEFwaS5HYW1lygIIQXBpXEdhbWXiAhRBcGlc",
            "R2FtZVxHUEJNZXRhZGF0YeoCCUFwaTo6R2FtZWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Api.Game.AdminHealthType), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum AdminHealthType {
    [pbr::OriginalName("AdminNone")] AdminNone = 0,
    [pbr::OriginalName("AdminSuccess")] AdminSuccess = 1,
  }

  #endregion

}

#endregion Designer generated code
