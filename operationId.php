<?php

$f = file_get_contents(getenv("WORKINGFILE"));

$j = json_decode($f);

foreach($j->paths as $k => $v) {
    if (str_contains($k, 'v2')) {
        if (property_exists($j->paths->$k, 'get') && property_exists($j->paths->$k->get, 'operationId')) {
            if (!str_contains('V2', $j->paths->$k->get->operationId)) {
                $j->paths->$k->get->operationId = $j->paths->$k->get->operationId.'V2';
            }
        }

        if (property_exists($j->paths->$k, 'post') && property_exists($j->paths->$k->post, 'operationId')) {
            if (!str_contains('V2', $j->paths->$k->post->operationId)) {
                $j->paths->$k->post->operationId = $j->paths->$k->post->operationId.'V2';
            }
        }

        if (property_exists($j->paths->$k, 'delete') && property_exists($j->paths->$k->delete, 'operationId')) {
            if (!str_contains('V2', $j->paths->$k->delete->operationId)) {
                $j->paths->$k->delete->operationId = $j->paths->$k->delete->operationId.'V2';
            }
        }
    }
    elseif (str_contains($k, 'v3') && str_contains($k, 'hebao')) {
        if (property_exists($j->paths->$k, 'get') && property_exists($j->paths->$k->get, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->get->operationId)) {
                $j->paths->$k->get->operationId = $j->paths->$k->get->operationId.'V3Hebao';
            }
        }

        if (property_exists($j->paths->$k, 'post') && property_exists($j->paths->$k->post, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->post->operationId)) {
                $j->paths->$k->post->operationId = $j->paths->$k->post->operationId.'V3Hebao';
            }
        }

        if (property_exists($j->paths->$k, 'delete') && property_exists($j->paths->$k->delete, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->delete->operationId)) {
                $j->paths->$k->delete->operationId = $j->paths->$k->delete->operationId.'V3Hebao';
            }
        }
    }
    elseif (str_contains($k, 'v3') && str_contains($k, 'amm')  && str_contains($k, 'poolsStats')) {
        if (property_exists($j->paths->$k, 'get') && property_exists($j->paths->$k->get, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->get->operationId)) {
                $j->paths->$k->get->operationId = $j->paths->$k->get->operationId.'AmmV3';
            }
        }

        if (property_exists($j->paths->$k, 'post') && property_exists($j->paths->$k->post, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->post->operationId)) {
                $j->paths->$k->post->operationId = $j->paths->$k->post->operationId.'AmmV3';
            }
        }

        if (property_exists($j->paths->$k, 'delete') && property_exists($j->paths->$k->delete, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->delete->operationId)) {
                $j->paths->$k->delete->operationId = $j->paths->$k->delete->operationId.'AmmV3';
            }
        }
    }
    elseif (str_contains($k, 'v3')) {
        if (property_exists($j->paths->$k, 'get') && property_exists($j->paths->$k->get, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->get->operationId)) {
                $j->paths->$k->get->operationId = $j->paths->$k->get->operationId.'V3';
            }
        }

        if (property_exists($j->paths->$k, 'post') && property_exists($j->paths->$k->post, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->post->operationId)) {
                $j->paths->$k->post->operationId = $j->paths->$k->post->operationId.'V3';
            }
        }

        if (property_exists($j->paths->$k, 'delete') && property_exists($j->paths->$k->delete, 'operationId')) {
            if (!str_contains('V3', $j->paths->$k->delete->operationId)) {
                $j->paths->$k->delete->operationId = $j->paths->$k->delete->operationId.'V3';
            }
        }
    }
}

file_put_contents(getenv("WORKINGFILE"), json_encode($j, JSON_UNESCAPED_SLASHES));
