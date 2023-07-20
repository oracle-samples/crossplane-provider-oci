#!/usr/bin/env bash

DEFAULT_SRG_EXAMPLES=compartment.yaml,bucket.yaml,vcn.yaml,subnet.yaml,instance.yaml
SRG_DATA_DIR=$(cd `dirname $0` && pwd)/testdata
SRG_FAILED=0

if [ "$1" = "--help" ]; then
  echo "Usage: ./srg.sh <optional ordered list of resource files under /testdata, or via evironment variable SRG_EXAMPLES>"
  echo "Default list of resource is $DEFAULT_SRG_EXAMPLES"
  exit 1
fi
if ! [ -z "$1" ]
  then
    SRG_EXAMPLES=$1
fi
if [ -z "$SRG_EXAMPLES" ]
  then
    SRG_EXAMPLES=$DEFAULT_SRG_EXAMPLES
fi
echo "=== Running SRG_EXAMPLES=$SRG_EXAMPLES ==="
echo "=== CREATE RESOURCES ==="
SRG_LOOP_STRING=$(echo $SRG_EXAMPLES | sed "s/,/ /g")
for fileName in $SRG_LOOP_STRING
do
  echo "kubectl apply -f $SRG_DATA_DIR/$fileName"
  kubectl apply -f $SRG_DATA_DIR/$fileName || (echo "FAILED applying $fileName, not ready"; exit 1)
done

echo "=== VERIFY CREATED RESOURCES ==="
for fileName in $SRG_LOOP_STRING
do
  echo "kubectl wait --for condition=ready -f $SRG_DATA_DIR/$fileName --timeout 180s"
  kubectl wait --for condition=ready -f $SRG_DATA_DIR/$fileName --timeout 180s || (echo "FAILED verifying $fileName, not ready"; SRG_FAILED=1; break)
done

echo "=== LIST RESOURCES ==="
echo "kubectl get managed"
kubectl get managed

# Delete resources in reversed order
echo "=== DELETE RESOURCES ==="
srg_array=($SRG_LOOP_STRING)
for i in `seq $((${#srg_array[@]} - 1)) -1 0`
 do
    fileName=${srg_array[$i]}
    echo "kubectl delete -f $SRG_DATA_DIR/$fileName --timeout 300s"
    kubectl delete -f "$SRG_DATA_DIR/$fileName" --timeout 300s || echo "FAILED deleting $fileName"
done

# Verify all resources deleted
echo "=== VERIFY DELETED RESOURCES ==="
echo "kubectl get managed"
(kubectl get managed  2>&1 | grep -v "crossplane-srg") || (echo "FAILED: some 'crossplane-srg' resources not deleted"; SRG_FAILED=1)

if [ $SRG_FAILED -eq 0 ];
  then
    echo "=== SRG PAASSED ==="
  else
    echo "=== SRG FAILED ==="
    exit 1
fi