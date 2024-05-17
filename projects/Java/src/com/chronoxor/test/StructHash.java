//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

package com.chronoxor.test;

public class StructHash implements Comparable<Object>
{
    public java.util.HashMap<String, Byte> f1 = new java.util.HashMap<String, Byte>();
    public java.util.HashMap<String, Byte> f2 = new java.util.HashMap<String, Byte>();
    public java.util.HashMap<String, java.nio.ByteBuffer> f3 = new java.util.HashMap<String, java.nio.ByteBuffer>();
    public java.util.HashMap<String, java.nio.ByteBuffer> f4 = new java.util.HashMap<String, java.nio.ByteBuffer>();
    public java.util.HashMap<String, EnumSimple> f5 = new java.util.HashMap<String, EnumSimple>();
    public java.util.HashMap<String, EnumSimple> f6 = new java.util.HashMap<String, EnumSimple>();
    public java.util.HashMap<String, FlagsSimple> f7 = new java.util.HashMap<String, FlagsSimple>();
    public java.util.HashMap<String, FlagsSimple> f8 = new java.util.HashMap<String, FlagsSimple>();
    public java.util.HashMap<String, StructSimple> f9 = new java.util.HashMap<String, StructSimple>();
    public java.util.HashMap<String, StructSimple> f10 = new java.util.HashMap<String, StructSimple>();

    public static final long fbeTypeConst = 141;
    public long fbeType() { return fbeTypeConst; }

    public StructHash() {}

    public StructHash(java.util.HashMap<String, Byte> f1, java.util.HashMap<String, Byte> f2, java.util.HashMap<String, java.nio.ByteBuffer> f3, java.util.HashMap<String, java.nio.ByteBuffer> f4, java.util.HashMap<String, EnumSimple> f5, java.util.HashMap<String, EnumSimple> f6, java.util.HashMap<String, FlagsSimple> f7, java.util.HashMap<String, FlagsSimple> f8, java.util.HashMap<String, StructSimple> f9, java.util.HashMap<String, StructSimple> f10)
    {
        this.f1 = f1;
        this.f2 = f2;
        this.f3 = f3;
        this.f4 = f4;
        this.f5 = f5;
        this.f6 = f6;
        this.f7 = f7;
        this.f8 = f8;
        this.f9 = f9;
        this.f10 = f10;
    }

    public StructHash(StructHash other)
    {
        this.f1 = other.f1;
        this.f2 = other.f2;
        this.f3 = other.f3;
        this.f4 = other.f4;
        this.f5 = other.f5;
        this.f6 = other.f6;
        this.f7 = other.f7;
        this.f8 = other.f8;
        this.f9 = other.f9;
        this.f10 = other.f10;
    }

    public StructHash clone()
    {
        // Serialize the struct to the FBE stream
        var writer = new com.chronoxor.test.fbe.StructHashModel();
        writer.serialize(this);

        // Deserialize the struct from the FBE stream
        var reader = new com.chronoxor.test.fbe.StructHashModel();
        reader.attach(writer.getBuffer());
        return reader.deserialize();
    }

    @Override
    public int compareTo(Object other)
    {
        if (other == null)
            return -1;

        if (!StructHash.class.isAssignableFrom(other.getClass()))
            return -1;

        final StructHash obj = (StructHash)other;

        int result = 0;
        return result;
    }

    @Override
    public boolean equals(Object other)
    {
        if (other == null)
            return false;

        if (!StructHash.class.isAssignableFrom(other.getClass()))
            return false;

        final StructHash obj = (StructHash)other;

        return true;
    }

    @Override
    public int hashCode()
    {
        int hash = 17;
        return hash;
    }

    @Override
    public String toString()
    {
        var sb = new StringBuilder();
        sb.append("StructHash(");
        if (f1 != null)
        {
            boolean first = true;
            sb.append("f1=[").append(f1.size()).append("][{");
            for (var item : f1.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                sb.append(item.getValue());
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append("f1=[0][{}]");
        if (f2 != null)
        {
            boolean first = true;
            sb.append(",f2=[").append(f2.size()).append("][{");
            for (var item : f2.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                if (item.getValue() != null) sb.append(item.getValue()); else sb.append("null");
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f2=[0][{}]");
        if (f3 != null)
        {
            boolean first = true;
            sb.append(",f3=[").append(f3.size()).append("][{");
            for (var item : f3.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                if (item.getValue() != null) sb.append("bytes[").append(item.getValue().array().length).append("]"); else sb.append("null");
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f3=[0][{}]");
        if (f4 != null)
        {
            boolean first = true;
            sb.append(",f4=[").append(f4.size()).append("][{");
            for (var item : f4.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                if (item.getValue() != null) sb.append("bytes[").append(item.getValue().array().length).append("]"); else sb.append("null");
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f4=[0][{}]");
        if (f5 != null)
        {
            boolean first = true;
            sb.append(",f5=[").append(f5.size()).append("][{");
            for (var item : f5.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                sb.append(item.getValue());
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f5=[0][{}]");
        if (f6 != null)
        {
            boolean first = true;
            sb.append(",f6=[").append(f6.size()).append("][{");
            for (var item : f6.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                if (item.getValue() != null) sb.append(item.getValue()); else sb.append("null");
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f6=[0][{}]");
        if (f7 != null)
        {
            boolean first = true;
            sb.append(",f7=[").append(f7.size()).append("][{");
            for (var item : f7.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                sb.append(item.getValue());
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f7=[0][{}]");
        if (f8 != null)
        {
            boolean first = true;
            sb.append(",f8=[").append(f8.size()).append("][{");
            for (var item : f8.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                if (item.getValue() != null) sb.append(item.getValue()); else sb.append("null");
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f8=[0][{}]");
        if (f9 != null)
        {
            boolean first = true;
            sb.append(",f9=[").append(f9.size()).append("][{");
            for (var item : f9.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                sb.append(item.getValue());
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f9=[0][{}]");
        if (f10 != null)
        {
            boolean first = true;
            sb.append(",f10=[").append(f10.size()).append("][{");
            for (var item : f10.entrySet())
            {
                if (item.getKey() != null) sb.append(first ? "" : ",").append("\"").append(item.getKey()).append("\""); else sb.append(first ? "" : ",").append("null");
                sb.append("->");
                if (item.getValue() != null) sb.append(item.getValue()); else sb.append("null");
                first = false;
            }
            sb.append("}]");
        }
        else
            sb.append(",f10=[0][{}]");
        sb.append(")");
        return sb.toString();
    }

    public String toJson() { return com.chronoxor.test.fbe.Json.getEngine().toJson(this); }
    public static StructHash fromJson(String json) { return com.chronoxor.test.fbe.Json.getEngine().fromJson(json, StructHash.class); }
}
